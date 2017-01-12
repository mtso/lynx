package lynx

import "testing"

func Test_titleFromFilename(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello-world.md", "Hello World"},
		{"sample.md", "Sample"},
		{"Sample.md", "Sample"},
		{"apples-plus-apples.mdown", "Apples Plus Apples"},
		{"A-(Code)-Walk-to-Remember.markdown", "A (Code) Walk to Remember"},
		{"sample.txt", "Sample"},
	}

	for _, c := range cases {
		got := titleFromFilename(c.in)
		if got != c.want {
			t.Errorf("func(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}