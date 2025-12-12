package main

import (
	"fmt"
	"log"
	"os"

	"github.com/voidpointer99/todocl/task"
	// "path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	manager, err := NewTaskManager()
	if err != nil {
		log.Fatalf("Failed to create the manager: %v", err)
	}

	switch command := os.Args[1]; command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo add <n> <description>")
			os.Exit(1)
		}
		todo := task.NewTask(os.Args[2], os.Args[3])
		if err := manager.Add(todo); err != nil {
			log.Fatalf("Failed to add task: %v", err)
		}
		fmt.Printf("✓ Task '%s' added successfully\n", todo.Name)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <n>")
			os.Exit(1)
		}
		if err := manager.Delete(os.Args[2]); err != nil {
			log.Fatalf("Failed to delete task: %v", err)
		}
		fmt.Printf("✓ Task '%s' deleted successfully\n", os.Args[2])

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <n>")
			os.Exit(1)
		}
		if err := manager.MarkDone(os.Args[2]); err != nil {
			log.Fatalf("Failed to mark task as done: %v", err)
		}
		fmt.Printf("✓ Task '%s' marked as done\n", os.Args[2])

	case "list":
		manager.List()

	case "help":
		printHelp()

	case "reset":
		if err := manager.Reset(); err != nil {
			log.Fatalf("Failed to reset tasks: %v", err)
		}
		fmt.Println("✓ All tasks cleared")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}
func printUsage() {
	fmt.Println("Usage: todo <command> [arguments]")
	fmt.Println("Run 'todo help' for more information")
}

func printHelp() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        📋 TODO CLI - HELP GUIDE                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	fmt.Println("📖 USAGE:")
	fmt.Println("   todo <command> [arguments]")
	fmt.Println()
	
	fmt.Println("✨ COMMANDS:")
	fmt.Println()
	
	// Add command
	fmt.Println("   \033[1madd\033[0m <n> <description>")
	fmt.Println("      Create a new task with a name and description")
	fmt.Println("      \033[2mExample: todo add \"Buy groceries\" \"Milk, eggs, bread\"\033[0m")
	fmt.Println()
	
	// List command
	fmt.Println("   \033[1mlist\033[0m")
	fmt.Println("      Display all tasks with their status and details")
	fmt.Println("      \033[2mExample: todo list\033[0m")
	fmt.Println()
	
	// Done command
	fmt.Println("   \033[1mdone\033[0m <n>")
	fmt.Println("      Mark a task as completed")
	fmt.Println("      \033[2mExample: todo done \"Buy groceries\"\033[0m")
	fmt.Println()
	
	// Delete command
	fmt.Println("   \033[1mdelete\033[0m <n>")
	fmt.Println("      Remove a task from your list")
	fmt.Println("      \033[2mExample: todo delete \"Buy groceries\"\033[0m")
	fmt.Println()
	
	// Reset command
	fmt.Println("   \033[1mreset\033[0m")
	fmt.Println("      Clear all tasks (use with caution!)")
	fmt.Println("      \033[2mExample: todo reset\033[0m")
	fmt.Println()
	
	// Help command
	fmt.Println("   \033[1mhelp\033[0m")
	fmt.Println("      Show this help message")
	fmt.Println("      \033[2mExample: todo help\033[0m")
	fmt.Println()
	
	fmt.Println("💡 TIPS:")
	fmt.Println("   • Task names must be unique")
	fmt.Println("   • Use quotes around names/descriptions with spaces")
	fmt.Println("   • Tasks are saved automatically to tasks.json")
	fmt.Println()
	
	fmt.Println("🚀 QUICK START:")
	fmt.Println("   1. Add your first task:    todo add \"Learn Go\" \"Complete the tutorial\"")
	fmt.Println("   2. View your tasks:        todo list")
	fmt.Println("   3. Mark it as done:        todo done \"Learn Go\"")
	fmt.Println()
}
