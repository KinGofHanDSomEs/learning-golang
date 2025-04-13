package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type WordCounter struct {
	wordsCount map[string]int
	mu         sync.Mutex
}

type CounterWorker interface {
	ProcessFiles(files ...string) error
	ProcessReader(r io.Reader) error
}

func (wc *WordCounter) ProcessFiles(files ...string) error {
	var wg sync.WaitGroup
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer file.Close()
			wc.ProcessReader(file)
		}()
	}
	wg.Wait()
	return nil
}

func (wc *WordCounter) ProcessReader(r io.Reader) error {
	scn := bufio.NewScanner(r)
	scn.Split(bufio.ScanWords)
	for scn.Scan() {
		wc.mu.Lock()
		wc.wordsCount[strings.ToLower(scn.Text())]++
		wc.mu.Unlock()
	}
	return nil
}

func main() {
	wordC := WordCounter{
		wordsCount: map[string]int{},
		mu:         sync.Mutex{},
	}
	wordC.ProcessFiles("file1", "file2")
	for key, val := range wordC.wordsCount {
		fmt.Printf("%v: %v\n", key, val)
	}
}
