package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/r4wm/walkUtils"
)

func main() {
	theDir := os.Args[1]

	err := filepath.Walk(theDir, walkUtils.PrintMD5Sum)

	if err != nil {
		log.Fatal(err)
	}

}
