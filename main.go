package main

import (
	"log"
	"os"

	"github.com/rohitpjpti18/tclient/file"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := file.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
