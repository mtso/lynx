package lynx

import "testing"

func TestSiteMapStruct (t *testing.T) {
	cases := []struct {
		in, want string
	}{
	}

	for _, c := range cases {
		got := ""//
		if got != c.want {
			t.Errorf("func(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}

func TestIsMarkdownFilename (t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"test.md", true},
		{"test.mdown", true},
		{"test.markdown", true},
		{"test.txt", false},
		{"test.html", false},
		{"test", false},
		{"test.", false},
	}

	for _, c := range cases {
		got := isMarkdownFilename(c.in)
		if got != c.want {
			t.Errorf("func(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}