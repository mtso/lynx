package main

import (
	// "fmt"
	"github.com/mtso/lynx"
)

func main() {
	// lynx.BuildPagesIn("content", "public")
	// lynx.TemplateDemo()
	// lynx.ProcessScss()
	pages := lynx.LoadPagesIn("content")
	// for _, p := range pages {
	// 	fmt.Println(p.Title, p.Next, p.Content)
	// }
	pages.Print()
}