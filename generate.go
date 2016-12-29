package lynx

import (
	"os"
)

func Generate() error {

	if err := mkdirIfNone("public"); err != nil {
		return err
	}

	pages, err := LoadPagesIn("content")
	if err != nil {
		return err
	}

	err = pages.ExportTo("public")
	if err != nil {
		return err
	}

	index := NewIndex("Blog", pages)
	index.loadTemplate("template/index-demo.html")
	index.executeTemplate()
	if err := index.writeTo("public"); err != nil {
		return err
	}

	return nil
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}