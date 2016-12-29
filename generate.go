package lynx

import (
	"log"
	"os"
)

func Generate() error {

	if err := mkdirIfNone("public"); err != nil {
		log.Fatal(err)
	}

	pages, err := lynx.LoadPagesIn("content")
	if err != nil {
		log.Fatal(err)
	}

	err = pages.ExportTo("public")
	if err != nil {
		log.Println(err)
	}

	index, err := lynx.NewIndex("Blog", pages)
	index.loadTemplate("template/index-demo.html")
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}