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

func assertSlicesEqual(got, want []byte, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("expected %q but got %q", want, got)
	} else {
		var areEqual = true;
		for i, wantItem := range want {
			if wantItem != got[i] {
				areEqual = false;
			}
		}
		if (!areEqual) {
			t.Errorf("expected %q but got %q", want, got)
		}
	}
	

}