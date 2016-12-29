package lynx

import (
	"os"
	"path/filepath"
)

const (
	indexTemplateName = "index.html"
)

func Generate() error {

	var err error = nil

	return CopyFromTo("./template/css/default.css", "./public/css/style.css")

	// Make public directory to store files
	if err := mkdirIfNone("public"); err != nil {
		return err
	}

	// Load content pages
	pages, err := LoadPagesIn("content")
	if err != nil {
		return err
	}

	path := filepath.Join("template", "post.html")
	if err := pages.loadTemplate(path); err != nil {
		return err
	}

	// Execute template
	pages.executeTemplate()

	// Save content pages
	err = pages.ExportTo("public")

	// Generate index page
	index := NewIndex("Blog", pages)
	indexTemplatePath := filepath.Join("template", indexTemplateName)
	if err := index.loadTemplate(indexTemplatePath); err != nil {
		return err
	}

	if err := index.executeTemplate(); err != nil {
		return err
	}

	if err := index.writeTo("public"); err != nil {
		return err
	}

	return err
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}

// // Copy css into public
// func copyCss() (err error) {
// 	destdir := "./public/css"
// 	srcdir := "./template/css"
// 	mkdirIfNone(destdir)
// 	dest := filepath.Join(destdir, "style.css")
// 	src := filepath.Join(srcdir, "default.css")
// 	os.Create(dest)
// 	err = CopyFile(dest, src)
// 	return
// }