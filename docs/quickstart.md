# Quick Start

Get gocmd2 running in 5 minutes.

## 1. Install

```bash
go get github.com/Necromancerlabs/gocmd2
```

## 2. Create a Shell

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

## 3. Run It

```bash
go run main.go
```

You now have a shell with `help`, `exit`, and module management commands built-in.

## 4. Add a Custom Module

```go
type MyModule struct {
    shell shellapi.ShellAPI
}

func (m *MyModule) Name() string { return "mymodule" }

func (m *MyModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
}

func (m *MyModule) GetCommands() []*cobra.Command {
    return []*cobra.Command{
        {
            Use:   "hello",
            Short: "Say hello",
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Hello, World!")
            },
        },
    }
}
```

Register it before `Run()`:

```go
sh.RegisterModule(&MyModule{})
sh.Run()
```

## Configuration

```go
// Change history file location
sh.SetHistoryFile("/path/to/history")

// Set custom prompt
sh.SetPrompt("myshell> ")

// Register exit handler
sh.OnExit(func() {
    fmt.Println("Goodbye!")
})
```

## Next Steps

- [Modules](modules.md) — Full module creation guide
- [Shell API](shell-api.md) — Complete API reference
- [Examples](examples.md) — More code examples
