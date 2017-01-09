package lynx

import "testing"

func Test_IsMarkdown(t *testing.T) {
	cases := []struct {
		in   filename
		want bool
	}{
		{filename("sample.md"), true},
		{filename("sample.markdown"), true},
		{filename("sample.mdown"), true},
		{filename("sample.html"), false},
		{filename("sample.txt"), false},
		{filename("sample"), false},
	}

	for _, c := range cases {
		got := c.in.isMarkdown()
		if got != c.want {
			t.Errorf("(`%q`).isMarkdown() == %q, expected %q", c.in, got, c.want)
		}
	}
}
