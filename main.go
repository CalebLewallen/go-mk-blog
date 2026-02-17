package main

import (
	"bufio"
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

//go:embed style.yaml home.md posts/* templates/* static/* robots.txt llm.txt
var content embed.FS

type StyleConfig struct {
	HeaderLogo      string
	Favicon         string
	BlogName        string
	PrimaryColor    string
	SecondaryColor  string
	LinkColor       string
	BackgroundColor string
}

type PostMetadata struct {
	PageName      string
	Published     bool
	PublishedDate time.Time
	Author        string
	RobotsAllowed bool
	Path          string // URL path
}

type Post struct {
	Metadata PostMetadata
	Content  template.HTML
}

var (
	config StyleConfig
	posts  []Post
	home   Post
)

func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := loadPosts(); err != nil {
		log.Fatalf("Error loading posts: %v", err)
	}

	http.HandleFunc("/posts", handlePostsList)
	http.HandleFunc("/static/", handleStatic)
	http.HandleFunc("/robots.txt", handleFile("robots.txt"))
	http.HandleFunc("/llm.txt", handleFile("llm.txt"))
	http.HandleFunc("/", handleRoot)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		handleHome(w, r)
		return
	}
	handlePost(w, r)
}

func render(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.ParseFS(content, "templates/layout.html", "templates/"+tmplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	latest := posts
	if len(latest) > 5 {
		latest = latest[:5]
	}
	data := struct {
		Title       string
		Config      StyleConfig
		Home        Post
		LatestPosts []Post
	}{
		Title:       "Home",
		Config:      config,
		Home:        home,
		LatestPosts: latest,
	}
	render(w, "home.html", data)
}

func handlePostsList(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Config StyleConfig
		Posts  []Post
	}{
		Title:  "All Posts",
		Config: config,
		Posts:  posts,
	}
	render(w, "posts.html", data)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var foundPost *Post
	for i := range posts {
		if posts[i].Metadata.Path == path {
			foundPost = &posts[i]
			break
		}
	}

	if foundPost == nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Title  string
		Config StyleConfig
		Post   *Post
	}{
		Title:  foundPost.Metadata.PageName,
		Config: config,
		Post:   foundPost,
	}
	render(w, "post.html", data)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(content)).ServeHTTP(w, r)
}

func handleFile(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	}
}

func loadConfig() error {
	data, err := content.ReadFile("style.yaml")
	if err != nil {
		return err
	}
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, "\"")

		switch key {
		case "headerLogo":
			config.HeaderLogo = val
		case "favicon":
			config.Favicon = val
		case "blogName":
			config.BlogName = val
		case "primaryColor":
			config.PrimaryColor = val
		case "secondaryColor":
			config.SecondaryColor = val
		case "linkColor":
			config.LinkColor = val
		case "backgroundColor":
			config.BackgroundColor = val
		}
	}
	return nil
}

func parseMarkdown(r io.Reader) (PostMetadata, string, error) {
	var meta PostMetadata
	var body strings.Builder
	scanner := bufio.NewScanner(r)

	if scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			for scanner.Scan() {
				line = scanner.Text()
				if line == "---" {
					break
				}
				parts := strings.SplitN(line, ":", 2)
				if len(parts) < 2 {
					continue
				}
				key := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])

				switch key {
				case "Page Name":
					meta.PageName = val
				case "Published":
					meta.Published = strings.ToLower(val) == "true"
				case "Published Date":
					t, _ := time.Parse("2006-01-02", val)
					meta.PublishedDate = t
				case "Author":
					meta.Author = val
				case "RobotsAllowed":
					meta.RobotsAllowed = strings.ToLower(val) == "true"
				}
			}
		} else {
			body.WriteString(line + "\n")
		}
	}

	for scanner.Scan() {
		body.WriteString(scanner.Text() + "\n")
	}

	return meta, body.String(), scanner.Err()
}

func renderMarkdown(md string) template.HTML {
	lines := strings.Split(md, "\n")
	var html strings.Builder

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			html.WriteString("<br>")
			continue
		}

		if strings.HasPrefix(trimmed, "# ") {
			html.WriteString(fmt.Sprintf("<h1>%s</h1>", line))
		} else if strings.HasPrefix(trimmed, "## ") {
			html.WriteString(fmt.Sprintf("<h2>%s</h2>", line))
		} else if strings.HasPrefix(trimmed, "### ") {
			html.WriteString(fmt.Sprintf("<h3>%s</h3>", line))
		} else if strings.HasPrefix(trimmed, "#### ") {
			html.WriteString(fmt.Sprintf("<h4>%s</h4>", line))
		} else if strings.HasPrefix(trimmed, "##### ") {
			html.WriteString(fmt.Sprintf("<h5>%s</h5>", line))
		} else if strings.HasPrefix(trimmed, "###### ") {
			html.WriteString(fmt.Sprintf("<h6>%s</h6>", line))
		} else if strings.HasPrefix(trimmed, "- ") {
			html.WriteString(fmt.Sprintf("<li>%s</li>", line))
		} else if strings.HasPrefix(trimmed, "* ") {
			html.WriteString(fmt.Sprintf("<li>%s</li>", line))
		} else {
			// Formatting (Bold, Italics, Links, Images)
			renderedLine := line

			// Images: ![alt](url "title") or ![alt](url) -> <img src="url" alt="alt" title="title">
			reImg := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+?)(?:\s+"([^"]+)")?\)`)
			renderedLine = reImg.ReplaceAllString(renderedLine, `<img src="$2" alt="$1" title="$3">`)

			// Links: [text](url) -> <a href="url">[text](url)</a>
			reLink := regexp.MustCompile(`(\[([^\]]+)\]\(([^)]+)\))`)
			renderedLine = reLink.ReplaceAllString(renderedLine, `<a href="$3">$1</a>`)

			// Bold: **text** -> <b>**text**</b>
			// We use [^*] to match characters between markers, but exclude the marker itself
			reBold := regexp.MustCompile(`(\*\*[^*]+\*\*)`)
			renderedLine = reBold.ReplaceAllString(renderedLine, "<b>$1</b>")

			// Italics: *text* -> <i>*text*</i>
			// Use a more careful regex to avoid matching inside <b> tags or existing stars
			// This is still basic but avoids the common double-matching of ** in *
			reItalic := regexp.MustCompile(`([^\*]|^)(\*[^*]+\*)([^\*]|$)`)
			renderedLine = reItalic.ReplaceAllString(renderedLine, "$1<i>$2</i>$3")

			html.WriteString(fmt.Sprintf("<p>%s</p>", renderedLine))
		}
	}

	return template.HTML(html.String())
}

func loadPosts() error {
	f, err := content.Open("home.md")
	if err != nil {
		return err
	}
	meta, body, err := parseMarkdown(f)
	f.Close()
	if err != nil {
		return err
	}
	home = Post{Metadata: meta, Content: renderMarkdown(body)}

	err = walkEmbedDir("posts", func(path string, d os.DirEntry) error {
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		f, err := content.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		meta, body, err := parseMarkdown(f)
		if err != nil {
			return err
		}
		if !meta.Published {
			return nil
		}
		relPath := strings.TrimPrefix(path, "posts/")
		relPath = strings.TrimSuffix(relPath, ".md")
		meta.Path = "/" + relPath
		posts = append(posts, Post{Metadata: meta, Content: renderMarkdown(body)})
		return nil
	})

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Metadata.PublishedDate.After(posts[j].Metadata.PublishedDate)
	})

	return err
}

func walkEmbedDir(dir string, fn func(path string, d os.DirEntry) error) error {
	entries, err := content.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if err := fn(path, entry); err != nil {
			return err
		}
		if entry.IsDir() {
			if err := walkEmbedDir(path, fn); err != nil {
				return err
			}
		}
	}
	return nil
}
