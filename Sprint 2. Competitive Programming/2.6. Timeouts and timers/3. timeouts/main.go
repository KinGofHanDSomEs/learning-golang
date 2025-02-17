package main

import (
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var result int
	for i, _ := range questions {
		timer := time.NewTimer(time.Second)
		select {
		case <-timer.C:
			continue
		case answer := <-answerCh:
			timer.Stop()
			if answer == answers[i] {
				result++
			}
		}
	}
	return result
}
