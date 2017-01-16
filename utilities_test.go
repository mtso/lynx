package lynx

import "testing"

func Test_stripExt(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"sample.md", "sample"},
		{"sample", "sample"},
		{"sample.template.html", "sample.template"},
	}

	for _, c := range cases {
		got := stripExt(c.in)
		if got != c.want {
			t.Errorf("func(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}

func Test_isMarkdownExtension(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"sample.md", true},
		{"sample.markdown", true},
		{"sample.mdown", true},
		{"sample.html", false},
		{"sample.txt", false},
		{"sample", false},
	}

	for _, c := range cases {
		got := isMarkdownExtension(c.in)
		if got != c.want {
			t.Errorf("isMarkdownExtension(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}
