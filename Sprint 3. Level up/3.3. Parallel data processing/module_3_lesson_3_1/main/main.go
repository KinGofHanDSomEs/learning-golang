package main

import (
	"os"
	"regexp"

	"hello/gzipper"
)

func main() {
	jobs := gzipper.FileNameGen(os.Args[1], regexp.MustCompile(`\.txt$`))
	gzipper.Compress(jobs)
}
