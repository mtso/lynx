package lynx

import (
	"testing"
	"os"
)

func Test_mkdirIfNone(t *testing.T) {
	testDir := "./test"
	_ = os.RemoveAll(testDir)
	err := mkdirIfNone(testDir)
	if err != nil {
		t.Errorf(err.Error())
	}
	assertExists(testDir, t)
}