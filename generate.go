package lynx

import (
	"os"
	"path/filepath"
)

const (
	indexTemplateName = "index.html"
	exportDir = "public"
)

func GenerateWith(config Configuration) (err error) {

	// Delete to refresh export folder
	err = os.RemoveAll(exportDir)

	// Make public directory to store files
	if err = mkdirIfNone(exportDir); err != nil {
		return
	}

	// Copy static files
	err = copyFromTo("./template/css/default.css", "./public/css/style.css")

	// Load content Pages
	Pages, err := loadPagesIn("content")
	if err != nil {
		return
	}

	path := filepath.Join("template", "post.html")
	if err = Pages.loadTemplate(path); err != nil {
		return
	}

	// Sort Pages in reverse chronological order
	Pages = Pages.reverseChronological()
	Pages.relinkNext()

	// Execute template
	Pages.executeTemplate()

	// Save content Pages
	err = Pages.exportTo(exportDir)

	// Generate index Page
	index := newIndex(config.Title, config.Description, Pages)
	indexTemplatePath := filepath.Join("template", indexTemplateName)
	if err = index.loadTemplate(indexTemplatePath); err != nil {
		return
	}

	if err = index.executeTemplate(); err != nil {
		return
	}

	if err = index.writeTo(exportDir); err != nil {
		return
	}

	return
}

func Generate() (err error) {
	return GenerateWith(Configuration{
		Title: "Blog",
		Description: "Blog description.",
	})
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}
