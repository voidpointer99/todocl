package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/voidpointer99/todocl/task"
)

type TaskManager struct {
	list []task.Task
}

func NewTaskManager() (*TaskManager, error) {
	manager := TaskManager{list: make([]task.Task, 0)}
	if err := manager.Load(); err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}
	return &manager, nil
}

func (t *TaskManager) Load() error {
	home, err := os.UserHomeDir()
	if err != nil { return err }
	path := filepath.Join(home, ".config", "todocl")
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("Failed to create The save folder %v error: %v", path, err)
	}
	path = filepath.Join(path, "tasks.json")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(path); err != nil {
			return fmt.Errorf("failed to create save file: %w", err)
		}
		return nil
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	if len(bytes) == 0 {
		return nil
	}

	if err := json.Unmarshal(bytes, &t.list); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func (t *TaskManager) Save() error {
	data, err := json.MarshalIndent(t.list, "", "  ")
	if err != nil {
		// CHANGED: Added error wrapping with %w
		return fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil { return err }
	path := filepath.Join(home, ".config", "todocl", "tasks.json") 
	if err := os.WriteFile(path, data, 0644); err != nil {
		// CHANGED: Added error wrapping with %w
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func (t *TaskManager) Add(task task.Task) error {
	// CHANGED: Added duplicate name validation
	// Prevents confusing situations where multiple tasks have the same name
	// Makes delete/done operations unambiguous
	for _, existing := range t.list {
		if existing.Name == task.Name {
			return fmt.Errorf("task with name '%s' already exists", task.Name)
		}
	}

	t.list = append(t.list, task)
	if err := t.Save(); err != nil {
		// CHANGED: More descriptive error message indicating when the error occurred
		return fmt.Errorf("failed to save after adding task: %w", err)
	}
	return nil
}

func (t *TaskManager) Delete(name string) error {
	for i, task := range t.list {
		if task.Name == name {
			t.list = append(t.list[:i], t.list[i+1:]...)
			if err := t.Save(); err != nil {
				// CHANGED: More descriptive error message
				return fmt.Errorf("failed to save after deletion: %w", err)
			}
			return nil
		}
	}
	// CHANGED: Now returns an error if task not found
	// Previously would silently succeed even if nothing was deleted
	// This gives user feedback that the task name was incorrect
	return fmt.Errorf("task '%s' not found", name)
}

// CHANGED: Added new MarkDone method to utilize the Done field
// This was missing functionality - the Done field existed but couldn't be set
func (t *TaskManager) MarkDone(name string) error {
	for i, task := range t.list {
		if task.Name == name {
			t.list[i].Done = true
			if err := t.Save(); err != nil {
				return fmt.Errorf("failed to save after marking done: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("task '%s' not found", name)
}

func (t *TaskManager) Reset() error {
	t.list = make([]task.Task, 0)
	if err := t.Save(); err != nil {
		// CHANGED: More descriptive error message
		return fmt.Errorf("failed to save after reset: %w", err)
	}
	return nil
}

func (t *TaskManager) List() {
	// CHANGED: Added check for empty list with user-friendly message
	// Better UX than printing nothing or just a header
	if len(t.list) == 0 {
		fmt.Println("\n📋 No tasks found. Add one with: todo add <name> <description>")
		return
	}

	// Count completed vs total tasks
	completedCount := 0
	for _, task := range t.list {
		if task.Done {
			completedCount++
		}
	}

	// CHANGED: Enhanced header with task statistics and visual separators
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║  📋 YOUR TASKS                                       %d/%d completed       ║\n", 
		completedCount, len(t.list))
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// CHANGED: Enhanced output format with better visual hierarchy
	// - Color-coded status indicators (emojis for better visibility)
	// - Aligned columns for cleaner look
	// - Dimmed completed tasks
	// - Time ago format instead of timestamp
	for i, task := range t.list {
		var statusIcon, taskDisplay string
		
		if task.Done {
			statusIcon = "✅"
			// Format completed tasks with strikethrough effect using unicode
			taskDisplay = fmt.Sprintf("\033[2m%s\033[0m", task.Name)
		} else {
			statusIcon = "⬜"
			taskDisplay = fmt.Sprintf("\033[1m%s\033[0m", task.Name)
		}

		// Calculate time ago
		timeAgo := formatTimeAgo(task.CreatedAt)
		
		fmt.Printf("  %s %2d. %s\n", statusIcon, i+1, taskDisplay)
		fmt.Printf("       %s\n", task.Description)
		fmt.Printf("       \033[2m⏱  Created %s\033[0m\n", timeAgo)
		
		// Add separator between tasks
		if i < len(t.list)-1 {
			fmt.Println("       ────────────────────────────────────────")
		}
	}
	
	fmt.Println()
}

// CHANGED: Added helper function to format timestamps as "time ago"
// More intuitive than showing exact timestamps (e.g., "2 hours ago" vs "2024-01-15 14:30")
func formatTimeAgo(t time.Time) string {
	duration := time.Since(t)
	
	if duration.Hours() < 1 {
		mins := int(duration.Minutes())
		if mins == 0 {
			return "just now"
		}
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	}
	
	if duration.Hours() < 24 {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}
	
	days := int(duration.Hours() / 24)
	if days == 1 {
		return "yesterday"
	}
	if days < 7 {
		return fmt.Sprintf("%d days ago", days)
	}
	if days < 30 {
		weeks := days / 7
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	}
	if days < 365 {
		months := days / 30
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	}
	
	return t.Format("Jan 2, 2006")
}
