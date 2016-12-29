package lynx

import (
	"strings"
)

type Filename string

func (f *Filename) isMarkdown() bool {
	raw := string(*f)
	switch {
	case strings.HasSuffix(raw, ".md"):
		fallthrough
	case strings.HasSuffix(raw, ".mdown"):
		fallthrough
	case strings.HasSuffix(raw, ".markdown"):
		return true
	default:
		return false
	}
}
