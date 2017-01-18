package main

import "testing"
import "log"

func Test_findFilenames(t *testing.T) {
	files, err := getFiles("./")
	if err != nil {
		t.Errorf("%q", err)
		return
	}
	for _, f := range files {
		log.Printf("found %q", f.Name())
	}
}
