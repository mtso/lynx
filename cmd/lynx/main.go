package main

import (
	"log"
	"github.com/mtso/lynx"
)

func main() {
	// lynx.BuildPagesIn("content", "public")
	// lynx.TemplateDemo()
	// lynx.ProcessScss()
	pages, err := lynx.LoadPagesIn("content")
	if err != nil {
		log.Fatal(err)
	}
	
	err = pages.ExportTo("public")
	if err != nil {
		log.Println(err)
	}
}