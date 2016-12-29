package lynx

import (
	"log"
	"io/ioutil"
	"os"
	"strings"
)

type SiteMap struct {
	Index string
	Pages []string
}

var (
	sections = []string{
		"head",
		"content",
		"footer",
	}
)

func (sm *SiteMap) LoadFilepathsIn(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		sm.Pages = append(sm.Pages, file.Name())
	}
	return nil
}

func LoadFilepathsIn(dirname string) {
	sm := &SiteMap{}
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		sm.Pages = append(sm.Pages, file.Name())
	}
}

func BuildPagesIn(dirname string, exportDir string) {
	sm := &SiteMap{}

	// Load filenames in dirname
	if err := sm.LoadFilepathsIn(dirname); err != nil {
		log.Fatal(err)
	}

	// Validate export directory
	if !strings.HasPrefix(exportDir, "./") {
		exportDir = "./" + exportDir
	}

	// Make export directory
	err := os.MkdirAll(exportDir, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		log.Fatal(err)
	}
	
	exportName := ""
	for _, filename := range sm.Pages {
		// load text
		buf, err := ioutil.ReadFile(dirname + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		// write html file
		exportName = exportDir + "/" + filename + ".html"
		_ = ioutil.WriteFile(exportName, buf, os.ModePerm)
	}

}

func isMarkdownExtension(filename string) bool {
	switch {
	case strings.HasSuffix(filename, ".md"):
		fallthrough
	case strings.HasSuffix(filename, ".mdown"):
		fallthrough
	case strings.HasSuffix(filename, ".markdown"):
		return true
	default:
		return false
	}
}