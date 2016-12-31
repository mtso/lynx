package lynx

import (
	"log"
	"path/filepath"
)

func stripExt(path string) string {
	i := 0
	// "012.45"
	// i = 3
	// path[:i] = 012
	for c, ch := range path {
		if ch == '.' {
			i = c 
		}
	}

	// if no extension existed
	// return original string 
	// no stripping necessary
	if i == 0 {
		return path
	}
	
	return path[:i]
}

// If the error is not nil
// logs it and returns true
func notNil(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func isMarkdownExtension(filename string) bool {
	switch {
	case filepath.Ext(filename) == ".md":
		return true
	case filepath.Ext(filename) == ".mdown":
		return true
	case filepath.Ext(filename) == ".markdown":
		return true
	default:
		return false
	}
}
