package lynx

import (
	"strings"
)

type Filename string

func (f *Filename) isMarkdown() bool {
	switch {
	case strings.HasSuffix(*f, ".md"):
		fallthrough
	case strings.HasSuffix(*f, ".mdown"):
		fallthrough
	case strings.HasSuffix(*f, ".markdown"):
		return true
	default:
		return false
	}
}
