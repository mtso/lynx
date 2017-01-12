package main

import (
	"log"
	"github.com/mtso/lynx"
)

func main() {
	
	if err := lynx.GenerateWith(lynx.Configuration{
		Title: "Blog Demo",
		Description: "This is a blog demo<br><a href=\"https://example.com\">hyperlink</a>",
	}); err != nil {
		log.Fatal(err)
	}
}