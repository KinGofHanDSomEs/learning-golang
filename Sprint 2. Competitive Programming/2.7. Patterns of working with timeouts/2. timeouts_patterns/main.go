package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	go func() {
		file, err := os.Open(path)
		if err != nil {
			result <- nil
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			result <- nil
			return
		}
		select {
		case <-ctx.Done():
			result <- nil
		case result <- data:
		}
	}()
}

func main() {
	result := make(chan []byte)
	ctx := context.Background()
	filePath := "data.json"
	go readJSON(ctx, filePath, result)
	select {
	case data := <-result:
		if data == nil {
			fmt.Print("Ошибка")
			return
		}
		fmt.Println(data)
	case <-time.After(2 * time.Second):
		fmt.Println("Превышено время ожидания")
	}
}
