package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/PuerkitoBio/goquery"
)

const GHOST_URL = "https://caleblewallen.com" // Change this to your actual Ghost URL if needed

type GhostExport struct {
	Db []Db `json:"db"`
}

type Db struct {
	Data Data `json:"data"`
}

type Data struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Title        string  `json:"title"`
	Slug         string  `json:"slug"`
	Html         *string `json:"html"`
	FeatureImage *string `json:"feature_image"`
	Type         string  `json:"type"`
	Status       string  `json:"status"`
	PublishedAt  *string `json:"published_at"`
}

func main() {
	jsonFile := "../caleb-lewallen.ghost.2026-02-22-19-09-53.json"

	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	var export GhostExport
	if err := json.Unmarshal(data, &export); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	if len(export.Db) == 0 {
		fmt.Println("No DB found in JSON")
		return
	}

	posts := export.Db[0].Data.Posts
	fmt.Printf("Found %d total items\n", len(posts))

	// Ensure directories exist
	os.MkdirAll("../posts", os.ModePerm)
	os.MkdirAll("../static/images", os.ModePerm)

	for _, post := range posts {
		if post.Type != "post" {
			continue
		}

		fmt.Printf("Processing post: %s\n", post.Title)

		// 1. Handle Feature Image
		var localFeatureImage string
		if post.FeatureImage != nil && *post.FeatureImage != "" {
			localPath, err := downloadImage(*post.FeatureImage)
			if err != nil {
				fmt.Printf("  Warning: Failed to download feature image %s: %v\n", *post.FeatureImage, err)
			} else {
				localFeatureImage = localPath
			}
		}

		// 2. Handle HTML and embedded images
		htmlContent := ""
		if post.Html != nil {
			htmlContent = *post.Html
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		if err != nil {
			fmt.Printf("  Error parsing HTML for post %s: %v\n", post.Slug, err)
			continue
		}

		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Attr("src")
			if exists {
				localPath, err := downloadImage(src)
				if err != nil {
					fmt.Printf("  Warning: Failed to download image %s: %v\n", src, err)
				} else {
					s.SetAttr("src", localPath)
					// Remove srcset and sizes to avoid mixed local/remote images
					s.RemoveAttr("srcset")
					s.RemoveAttr("sizes")
				}
			}
		})

		modifiedHtml, err := doc.Find("body").Html()
		if err != nil {
			modifiedHtml = htmlContent // fallback
		}

		// 3. Convert HTML to Markdown
		mdContent, err := htmltomarkdown.ConvertString(modifiedHtml)
		if err != nil {
			fmt.Printf("  Error converting HTML to Markdown for post %s: %v\n", post.Slug, err)
			mdContent = modifiedHtml // fallback to HTML if conversion fails
		}

		// 4. Prepare Frontmatter
		published := "false"
		if post.Status == "published" {
			published = "true"
		}

		publishedDate := ""
		if post.PublishedAt != nil {
			publishedDate = *post.PublishedAt
		}

		var buf bytes.Buffer
		buf.WriteString("---\n")
		buf.WriteString(fmt.Sprintf("Page Name: %s\n", post.Title))
		buf.WriteString(fmt.Sprintf("Published: %s\n", published))
		buf.WriteString(fmt.Sprintf("Published Date: %s\n", publishedDate))
		buf.WriteString("Author: Caleb Lewallen\n")
		buf.WriteString("RobotsAllowed: true\n")
		buf.WriteString("---\n\n")

		if localFeatureImage != "" {
			buf.WriteString(fmt.Sprintf("![Feature Image](%s)\n\n", localFeatureImage))
		}

		buf.WriteString(mdContent)
		buf.WriteString("\n")

		// 5. Write to file
		mdFilePath := filepath.Join("../posts", post.Slug+".md")
		err = os.WriteFile(mdFilePath, buf.Bytes(), 0644)
		if err != nil {
			fmt.Printf("  Error writing markdown file %s: %v\n", mdFilePath, err)
		} else {
			fmt.Printf("  Created %s\n", mdFilePath)
		}
	}
	fmt.Println("Migration complete!")
}

func downloadImage(imgURL string) (string, error) {
	// Replace __GHOST_URL__ if present
	imgURL = strings.ReplaceAll(imgURL, "__GHOST_URL__", GHOST_URL)

	// Handle relative URLs
	if strings.HasPrefix(imgURL, "/") {
		imgURL = strings.TrimSuffix(GHOST_URL, "/") + imgURL
	}

	parsedURL, err := url.Parse(imgURL)
	if err != nil {
		return "", err
	}

	// Extract filename
	filename := path.Base(parsedURL.Path)
	if filename == "" || filename == "/" {
		filename = fmt.Sprintf("image-%d.jpg", time.Now().UnixNano())
	}

	// Ensure unique filename if it already exists (simple approach)
	localPath := filepath.Join("../static/images", filename)
	if _, err := os.Stat(localPath); err == nil {
		// File exists, append timestamp
		ext := filepath.Ext(filename)
		name := strings.TrimSuffix(filename, ext)
		filename = fmt.Sprintf("%s-%d%s", name, time.Now().Unix(), ext)
		localPath = filepath.Join("../static/images", filename)
	}

	// Download the file
	resp, err := http.Get(imgURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(localPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	// Return the absolute path from the root of the site
	return "/static/images/" + filename, nil
}
