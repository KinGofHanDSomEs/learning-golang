package gzipper

import (
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Work struct {
	FilePath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work)
	go func() {
		defer close(jobs)

		filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "input.txt" || info.Name() == "output.txt" {
				return nil
			}
			if !info.IsDir() && pattern.MatchString(info.Name()) {
				jobs <- Work{FilePath: path}
			}
			return nil
		})
	}()
	return jobs
}

func Compress(jobs <-chan Work) {
	var wg sync.WaitGroup
	for work := range jobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inputFile, err := os.Open(work.FilePath)
			if err != nil {
				return
			}
			defer inputFile.Close()
			outputFile, err := os.Create(work.FilePath + ".gz")
			if err != nil {
				return
			}
			defer outputFile.Close()
			gzipWriter := gzip.NewWriter(outputFile)
			defer gzipWriter.Close()
			io.Copy(gzipWriter, inputFile)
		}()
	}
	wg.Wait()
}
