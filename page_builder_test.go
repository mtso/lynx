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