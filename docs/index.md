---
layout: default
title: Home
nav_order: 1
---

# gocmd2

A Go-based interactive shell framework inspired by Python's [cmd2](https://github.com/python-cmd2/cmd2) library.

gocmd2 provides a modular, extensible framework for building interactive command-line applications with rich features like command auto-completion, history management, and a module-based architecture.

---

## Features

| Feature | Description |
|---------|-------------|
| **Modular Architecture** | Create and register command modules to organize functionality |
| **Built on Cobra** | Uses [Spf13 Cobra](https://github.com/spf13/cobra) for robust command parsing |
| **Tab Completion** | Command auto-completion via [Readline](https://github.com/chzyer/readline) |
| **Command History** | Persistent command history between sessions |
| **Runtime Module Control** | Enable/disable modules at runtime |
| **Shared State** | Thread-safe state sharing between modules |
| **Exit Handlers** | Register cleanup functions on shell exit |

---

## Quick Example

```go
package main

import (
    "fmt"
    "os"

    "github.com/Necromancerlabs/gocmd2/pkg/shell"
)

func main() {
    sh, err := shell.NewShell("myshell", "Welcome!")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
    defer sh.Close()

    sh.Run()
}
```

---

## Documentation

- [Getting Started](getting-started) - Installation and basic usage
- [Modules](modules) - Creating and managing command modules
- [Shell API](shell-api) - Complete API reference
- [Core Commands](core-commands) - Built-in commands reference
- [Examples](examples) - Code examples and patterns

---

## Installation

```bash
go get github.com/Necromancerlabs/gocmd2
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| [spf13/cobra](https://github.com/spf13/cobra) | Command framework |
| [chzyer/readline](https://github.com/chzyer/readline) | Interactive input with history |
| [google/shlex](https://github.com/google/shlex) | Shell-like argument parsing |

---

## License

See the [LICENSE](https://github.com/Necromancer-Labs/gocmd2/blob/main/LICENSE) file for details.
