package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskManager struct {
	list 		[]Task
}

func (t *TaskManager) Load() error {
	if _, err := os.Stat(SAVE_FILE); os.IsNotExist(err) {
		if _, err := os.Create(SAVE_FILE); err != nil {
			return err 
		}
		return nil	
	}
	
	bytes, err := os.ReadFile(SAVE_FILE)
	if err != nil {
		return fmt.Errorf("Failed to read the file: %s", SAVE_FILE)
	}
	if len(bytes) == 0 {
		return  nil
	}
	err = json.Unmarshal(bytes, &t.list)
	if err != nil {
		return  fmt.Errorf("Failed To Unmarshal The Json")
	}
	return nil
}

func (t *TaskManager) Reset() error {
	t.list = make([]Task, 0)
	if err := t.Save(); err != nil {
		return fmt.Errorf("Failed to Save the Cache: %w", err)
	}
	return nil
}

func NewTaskManager() (*TaskManager, error) {
	manager := TaskManager {list: make([]Task, 0)}

	err := manager.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to Load the Cache: %w", err)
	}
	return &manager, nil
}

func (t *TaskManager) Add(task Task) error {
	t.list = append(t.list, task)
	if err := t.Save(); err != nil {
		return fmt.Errorf("Failed to Save the Cache: %w", err)
	}
	return nil
}
func (t *TaskManager) Delete(name string) error {
	deleted := false
	for i, task := range t.list {
		if task.Name == name {
			t.list = append(t.list[:i], t.list[i+1:]...)
			deleted = true
			break
		}
	}
	if deleted {
		if err := t.Save(); err != nil {
			return fmt.Errorf("Failed to Save the deletion")
		}
	}
	return nil
}

func (t *TaskManager) Save() error {
	data, err := json.Marshal(t.list)
	if err != nil {
		return fmt.Errorf("Failed TO Marshal to Json")
	}
	err = os.WriteFile(SAVE_FILE, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed TO Write To the File")
	}
	return nil
}

func (t *TaskManager) List() {
	for _, task := range t.list {
		fmt.Println("", task.Name, task.Description)
	}
}


