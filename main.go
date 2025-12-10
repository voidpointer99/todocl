package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const SAVE_FILE string = "tasks.json"

func main() {
	manager, err := NewTaskManager() 
	if err != nil {
		log.Fatal("Failed to Create The manager: %w", err)
	}
	
	switch firstArg := os.Args[1]; firstArg {
		case "add":
			if len(os.Args) < 4 {
				fmt.Println("Usage: todo add <name> <description>")
				os.Exit(1)
			}
			task := NewTask(os.Args[2], os.Args[3]) 
			manager.Add(task)
		case "delete":
			manager.Delete(os.Args[2])
		case "list":
			manager.List()
		case "reset":
			manager.Reset()
	}

}
