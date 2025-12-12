package task

import "time"


type Task struct {
	Name 						string 			
	Description 		string 		
	Done						bool 				
	CreatedAt				time.Time
}
func NewTask(name string, desc string) Task {
	return Task{Name: name, Description: desc, Done: false, CreatedAt: time.Now()}
}


