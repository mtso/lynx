package lynx

import "testing"


func Test_IsMarkdown (t *testing.T) {
	cases := []struct {
		in Filename
		want bool
	}{
		{Filename("sample.md"), true},
		{Filename("sample.markdown"), true},
		{Filename("sample.mdown"), true},
		{Filename("sample.html"), false},
		{Filename("sample.txt"), false},
		{Filename("sample"), false},
	}

	for _, c := range cases {
		got := c.in.isMarkdown()
		if got != c.want {
			t.Errorf("(`%q`).isMarkdown() == %q, expected %q", c.in, got, c.want)
		}
	}
}

