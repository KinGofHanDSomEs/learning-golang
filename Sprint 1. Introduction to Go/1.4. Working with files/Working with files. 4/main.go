package main

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	file, _ := os.OpenFile(filename, os.O_WRONLY, 0600)
	defer file.Close()
	file.Seek(int64(pos), 0)
	file.WriteString(val)
}