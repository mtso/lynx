package lynx

import (
	"fmt"
)

var (
	sections = [...]string{
		"head",
		"content",
		"footer",
	}
)

func PrintSections() {
	for _, s := range sections {
		fmt.Println(s)
	}
}
