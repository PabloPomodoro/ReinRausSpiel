package main

import "time"

type item struct {
	name    string
	dueDate time.Time
	done    bool
}

type todoList struct {
	items []item
}
