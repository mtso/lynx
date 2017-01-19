package lynx

import "testing"
import "time"
import "log"

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


	if gotTitle != wantTitle {
		t.Errorf("New page title: %q, but expected %q", gotTitle, wantTitle)
	}

	// expectbuffer := []byte("<p>test data</p>")
	// writebuffer := []byte("<p>test data</p>")
	// writelen, _ := testPage.Write(writebuffer)
	// log.Println(testPage.html)

	// readbuffer := make([]byte, writelen)
	// readlen, _ := testPage.Read(readbuffer)
	// log.Println(readlen)
	// log.Println(readbuffer)
	// log.Println(testPage.html)
	// assertSlicesEqual(readbuffer, expectbuffer, t)
}

func Test_timeComparisons(t *testing.T) {
	pageBefore := newPage("PageBefore", nil, "test", time.Now(), "./test-title", "<h1>Heading 1</h1>", time.Now(), "10")
	time.Sleep(100)
	pageAfter := newPage("PageAfter", nil, "test", time.Now(), "./test-title", "<h1>Heading 1</h1>", time.Now(), "10")

	cases := []struct {
		before, after *Page
		actual, want  bool
	}{
		{pageBefore, pageAfter, pageBefore.isModifiedBefore(*pageAfter), true},
		{pageBefore, pageAfter, pageAfter.isModifiedAfter(*pageBefore), true},
		{pageBefore, pageAfter, pageBefore.isCreatedBefore(*pageAfter), true},
		{pageBefore, pageAfter, pageAfter.isCreatedAfter(*pageBefore), true},
		{pageBefore, pageAfter, pageAfter.isModifiedBefore(*pageBefore), false},
		{pageBefore, pageAfter, pageBefore.isModifiedAfter(*pageAfter), false},
		{pageBefore, pageAfter, pageAfter.isCreatedBefore(*pageBefore), false},
		{pageBefore, pageAfter, pageBefore.isCreatedAfter(*pageAfter), false},
	}

	for _, c := range cases {
		if c.actual != c.want {
			t.Errorf("%q.isModifiedBefore(%q) Expected %v, but got %v", c.before, c.after, c.want, c.actual)
		}
	}
}
