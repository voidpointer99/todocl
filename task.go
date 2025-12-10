package main

import "time"


type Task struct {
	Name 						string 		`json:"name"`	
	Description 		string 		`json:"description"`
	Done						bool 			`json:"done"`
	CreatedAt				time.Time `json:"createdat"`
}
func NewTask(name string, desc string) Task {
	return Task{Name: name, Description: desc, Done: false, CreatedAt: time.Now()}
}


