package main

import (
	"sync"

	"github.com/kingofhandsomes/game_of_life/application"
	"github.com/kingofhandsomes/game_of_life/game"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		application.RunServer("8081")
	}()
	go func() {
		defer wg.Done()
		game.RunGame()
	}()
	wg.Wait()
}
