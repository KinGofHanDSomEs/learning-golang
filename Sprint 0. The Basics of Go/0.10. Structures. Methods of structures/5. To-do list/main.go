package main

import "time"

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() bool {
	return t.priority >= 4
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {
	return len(t.notes)
}

func (t ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsTopPriority() {
			count++
		}
	}
	return count
}

func (t ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}
