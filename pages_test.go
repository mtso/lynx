package lynx

import "testing"

func Test_pageOperations(t *testing.T) {

	pages, err := loadPagesIn("content")
	if err != nil {
		t.Errorf("error loading from content dir: %s", err)
	}

	firstPage := &pages[0]
	pages.reverse()
	if firstPage.Title != pages[0].Title {
		t.Errorf("FAILED [reverses pages array order] expected %q but got %q", firstPage.Title, pages[0].Title)
	}

	pages.chronological()
	lastPage := &pages[len(pages)-1]
	if !firstPage.isCreatedBefore(*lastPage) {
		t.Errorf("FAILED [sorts pages chronologically] expected %q to be before %q", firstPage.BirthTime, lastPage.BirthTime)
	}
}
