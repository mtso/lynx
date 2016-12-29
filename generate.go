package lynx

import (
	"os"
	"path/filepath"
)

const (
	indexTemplateName = "index.html"
	exportDir = "public"
)

func Generate() (err error) {

	// Delete to refresh export folder
	err = os.RemoveAll(exportDir)

	// Make public directory to store files
	if err = mkdirIfNone(exportDir); err != nil {
		return
	}

	// Copy static files
	err = CopyFromTo("./template/css/default.css", "./public/css/style.css")

	// Load content pages
	pages, err := LoadPagesIn("content")
	if err != nil {
		return
	}

	path := filepath.Join("template", "post.html")
	if err = pages.loadTemplate(path); err != nil {
		return
	}

	// Execute template
	pages.executeTemplate()

	// Save content pages
	err = pages.ExportTo(exportDir)

	// Generate index page
	index := NewIndex("Blog", pages)
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

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}
