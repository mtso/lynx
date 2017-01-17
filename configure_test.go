package lynx

import "testing"

func Test_ConfigurationStruct(t *testing.T) {
	cases := []struct {
		inTitle, inDesc, wantTitle, wantDesc string
	}{
		{"Test Title", "Test Description", "Test Title", "Test Description"},
	}

	for _, c := range cases {
		conf := &Configuration{c.inTitle, c.inDesc}
		gotTitle := conf.Title
		gotDesc := conf.Description
		if (conf.Title != c.wantTitle && conf.Description != c.wantDesc) {
			t.Errorf("New Configuration title/desc: Expected %q/%q, got %q/%q", c.wantTitle, c.wantDesc, gotTitle, gotDesc)
		}
	}
}