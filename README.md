# 📋 Go Task Manager CLI

A simple, fast, and persistent command-line To-Do list manager built with Go.

This CLI tool allows you to easily manage tasks directly from your terminal, saving your list to a JSON file for persistence across sessions.

## ✨ Features

* **Persistent Storage:** Tasks are automatically saved to `tasks.json` in your user's config directory (`~/.config/todocl/`).
* **Simple CRUD Operations:** Easily add, list, mark done, and delete tasks.
* **Enhanced Listing:** Displays task status (completed/pending), descriptions, and creation time formatted as "time ago" for better readability.
* **Robust Error Handling:** Provides clear feedback and error messages for invalid commands or non-existent tasks.

## 🚀 Installation

### Prerequisites

You must have [Go installed](https://go.dev/doc/install) on your system.

### Method 1: Using `go install` (Recommended)

Run the following command in your terminal. This will compile the code and place the executable (`todo` or `todo.exe`) in your `$GOPATH/bin` directory, making it globally accessible.

```bash
go install [github.com/voidpointer99/todocl@latest](https://github.com/voidpointer99/todocl@latest)
