package lynx

import (
	"testing"
	"os"	
)

func assertExists(filepath string, t *testing.T) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("Expected %q to exist, but it does not: %q", filepath, err.Error());
	}
}

func Test_testAssert(t *testing.T) {
	var filepath = "./test"
	err := mkdirIfNone(filepath)
	if err != nil {
		t.Errorf(err.Error())
	}
	assertExists(filepath, t)
}