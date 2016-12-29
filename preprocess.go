package lynx

import (
	"bytes"
	"os"
	"log"
	"github.com/wellington/go-libsass"
)

func ProcessScss() {
	buf := bytes.NewBufferString("div { p { color: red; } }")
	css, err := libsass.New(os.Stdout, buf)
	if err != nil {
		log.Fatal(err)
	}
	if err := css.Run(); err != nil {
		log.Fatal(err)
	}
}
