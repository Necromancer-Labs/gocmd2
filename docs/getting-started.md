---
layout: default
title: Getting Started
nav_order: 2
---

# Getting Started

This guide will help you set up gocmd2 and create your first interactive shell.

---

## Installation

```bash
go get github.com/Necromancerlabs/gocmd2
```

---

## Creating a Basic Shell

The simplest shell requires only a few lines of code:

```go
package main

import (
    "fmt"
    "os"

    "github.com/Necromancerlabs/gocmd2/pkg/shell"
)

func main() {
    // Create a new shell with name and banner
    sh, err := shell.NewShell("myshell", "Welcome to My Shell!")
    if err != nil {
        fmt.Printf("Error initializing shell: %v\n", err)
        os.Exit(1)
    }
    defer sh.Close()

    // Run the interactive shell
    sh.Run()
}
```

This creates a shell with:
- Core commands (`help`, `exit`, `modules`, `enable`, `disable`)
- Tab completion
- Command history
- A customizable prompt

---

## Shell Configuration

### Constructor Parameters

```go
shell.NewShell(name string, banner string) (*Shell, error)
```

| Parameter | Description |
|-----------|-------------|
| `name` | The root command name (appears in help) |
| `banner` | Welcome message displayed on shell start |

### History File

By default, command history is stored in `/tmp/readline.tmp`. Change it with:

```go
sh.SetHistoryFile("/path/to/history")
```

### Custom Prompt

Set the shell prompt dynamically:

```go
sh.SetPrompt("myshell> ")
```

---

## Adding Exit Handlers

Register cleanup functions that run when the shell exits:

```go
sh.OnExit(func() {
    fmt.Println("Cleaning up resources...")
    // Your cleanup code here
})
```

Multiple handlers can be registered and will execute in order.

---

## Running Commands Programmatically

Execute commands without user input:

```go
sh.ExecuteCommand("help")
```

This is useful for:
- Testing
- Scripted automation
- Initial setup commands

---

## Project Structure

A typical gocmd2 project structure:

```
myproject/
├── main.go           # Shell initialization
├── modules/
│   ├── mymodule.go   # Custom module
│   └── another.go    # Another module
├── go.mod
└── go.sum
```

---

## Next Steps

- Learn how to [create custom modules](modules)
- Explore the [Shell API](shell-api)
- See [example implementations](examples)
