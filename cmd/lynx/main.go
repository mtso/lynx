package main

import (
	"log"
	"github.com/mtso/lynx"
)

func main() {
	
	if err := lynx.Generate(); err != nil {
		log.Fatal(err)
	}
}