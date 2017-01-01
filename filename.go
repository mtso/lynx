package lynx

import (
	"strings"
)

type filename string

func (f *filename) isMarkdown() bool {
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
