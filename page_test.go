package lynx

import "testing"
import "time"

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
		{"sample", "Sample"},
		{"sample.html", "Sample"},
		{"Hello World.md", "Hello World"},
	}

	for _, c := range cases {
		got := titleFromFilename(c.in)
		if got != c.want {
			t.Errorf("func(%q) == %q, expected %q", c.in, got, c.want)
		}
	}
}

func Test_newPage(t *testing.T) {
	testPage := newPage("Test Title", nil, "test", time.Now(), "./test-title", "<h1>Heading 1</h1>", time.Now(), "10")
	gotTitle := testPage.String()
	wantTitle := "Test Title"

	if (gotTitle != wantTitle) {
		t.Errorf("New page title: %q, but expected %q", gotTitle, wantTitle)
	}
}