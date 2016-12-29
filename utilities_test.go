package lynx

import "testing"

func Test_StripExt(t *testing.T) {
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
