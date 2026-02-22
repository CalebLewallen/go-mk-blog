# Caleb Lewallen Personal Blog

A lightweight, fast, and simple personal blog built with Go. This project was created to explore Go's `embed` package and build a dependency-free blog engine that serves Markdown files directly from the binary.

## Features

- **Markdown-based:** Write posts in plain Markdown with simple frontmatter.
- **Embedded Assets:** Uses Go's `embed` package to bundle templates, static files, and markdown content directly into the compiled binary.
- **Zero-Database:** All content is served from the embedded file system.
- **Customizable Styling:** Configured via a simple `style.yaml` file.
- **Built-in Routing:** Directory-based routing for posts, along with `/posts` for a full list and `/` for the home page.
- **SEO & AI Friendly:** Serves `robots.txt` and `llm.txt` out of the box.
- **Custom Markdown Parser:** A lightweight, built-in markdown parser that handles headers, lists, links, images, bold, and italics without external dependencies.

## Project Structure

- `main.go`: The core web server and markdown parser.
- `posts/`: Directory containing all blog posts in Markdown format.
- `templates/`: HTML templates (`layout.html`, `home.html`, `post.html`, `posts.html`) using Go's standard `html/template`.
- `static/`: Static assets like CSS, images, and fonts.
- `style.yaml`: Configuration for blog name, colors, and theme.
- `home.md`: The markdown content for the home page.

## Running Locally

It's incredibly simple to run. Just clone the repository and start the Go server:

```bash
git clone https://github.com/CalebLewallen/go-mk-blog.git
cd go-mk-blog
go run main.go
```

The server will start on `http://localhost:8080`.

## Writing Posts

To create a new post, simply add a new `.md` file to the `posts/` directory. The server parses metadata from the top of the markdown files.

Example post format:

```markdown
---
Page Name: My Awesome Post
Published: true
Published Date: 2026-02-22
Author: Caleb Lewallen
RobotsAllowed: true
---

# Post Title

Your markdown content goes here...
```

## Configuration

You can customize the look and feel of the blog by editing `style.yaml`:

```yaml
headerLogo: "/static/images/caleblewallen-website-logo.png"
favicon: "/static/images/fibonacci.svg"
blogName: "Caleb Lewallen"
primaryColor: "#265D4F"
secondaryColor: "#d8280a"
linkColor: "#265D4F"
backgroundColor: "#fcf9ec"
theme: "terminal"
```