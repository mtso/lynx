package lynx

import (
	"bytes"
	"github.com/wellington/go-libsass"
	"log"
	"os"
	"strings"
	"io/ioutil"
	"path"
)

func processScss(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !strings.HasPrefix(f.Name(), "_") && path.Ext(f.Name()) == ".scss" {
			buf, err := ioutil.ReadFile(path.Join(dirname, f.Name()))
			b := bytes.NewBufferString(string(buf))
			if err != nil {
				log.Fatal(err)
			}
			css, err := libsass.New(os.Stdout, b)
			if err != nil {
				log.Fatal(err)
			}
			if err := css.Run(); err != nil {
				log.Fatal(err)
			}
		}
	}

	// buf := bytes.NewBufferString(`$heading1: 4em; h1 { font-size: $heading1; } div { p { color: red; } }`)
	// css, err := libsass.New(os.Stdout, buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := css.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
