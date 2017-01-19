package lynx

import (
	"testing"
	"os"
)

func Test_generateDefault(t *testing.T) {
	testDir := defaultExportDir
	Generate()
	assertExists(testDir, t)
	_ = os.RemoveAll(testDir)
}

func Test_generateWith(t *testing.T) {
	testDir := "./generate_test"
	testconfig := &Configuration{
		Title: "Generation Test Title",
		Description: "Generation Test Description",
		ExportDir: &testDir,
	}
	GenerateWith(*testconfig)
	assertExists(testDir, t)
	_ = os.RemoveAll(testDir)
}

func Test_mkdirIfNone(t *testing.T) {
	testDir := "./test"
	err := mkdirIfNone(testDir)
	if err != nil {
		t.Errorf(err.Error())
	}
	assertExists(testDir, t)

	// Clean test directories
	_ = os.RemoveAll(testDir)
}