package lynx

import (
	
)

type Index struct {

	Title string

	Pages Pages
}

func NewIndex(t string, pgs Pages) *Index {
	return &Index {
		Title: t,
		Pages: pgs,
	}
}