package lynx

import (
	"testing"
	"time"
)

// Function prototypes to test:
// func parseDateFrom(raw string) (time.Time, error)
// func parseFrontMatterIn(b []byte) (map[string]interface{}, error)
// func parseFrontMatterLine(line string) (string, interface{}, error)
// func stripFrontMatterFrom(b []byte) []byte

func Test_parseDateFrom(t *testing.T) {
	cases := []struct {
		input string
		date  time.Time
		err   error
	}{
		{"jan 7 2017", time.Date(2017, time.January, 7, 0, 0, 0, 0, time.UTC), nil},
		{"Jan 7 2017", time.Date(2017, time.January, 7, 0, 0, 0, 0, time.UTC), nil},
		{"jan 7, 2017", time.Time{}, &time.ParseError{}}, //Message: "Invalid Format"}},
	}

	for _, c := range cases {
		got, err := parseDateFrom(c.input)

		isSameYear := got.Year() == c.date.Year()
		isSameMonth := got.Month() == c.date.Month()
		isSameDay := got.Day() == c.date.Day()

		if (!isSameYear || !isSameMonth || !isSameDay) && err == nil {
			t.Errorf("parseDateFrom(%q) == %q, %q; expected %q, %q", c.input, got, err, c.date, c.err)
		}
	}
}

func Test_stripsFrontMatter(t *testing.T) {
	cases := []struct {
		in, want []byte
	}{
		{[]byte("---\ndate: jan 21 2017\n---\n<DOCTYPE! html>\n<html>\n<head>\n<title>blog</title>\n</head>\n</html>"), []byte("<DOCTYPE! html>\n<html>\n<head>\n<title>blog</title>\n</head>\n</html>")},
		{[]byte("<DOCTYPE! html>\n<html>\n<head>\n<title>blog</title>\n</head>\n</html>"), []byte("<DOCTYPE! html>\n<html>\n<head>\n<title>blog</title>\n</head>\n</html>")},
	}

	for _, c := range cases {
		got := stripFrontMatterFrom(c.in)

		if string(got) != string(c.want) {
			t.Errorf("stripsFrontMatterFrom(), got: %q", got)
		}
	}
}
