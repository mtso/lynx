package lynx

import(
	
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
	if i == 0 {
		return path
	}
	return path[:i]
}