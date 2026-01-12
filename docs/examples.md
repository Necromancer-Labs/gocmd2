# Examples

Practical examples demonstrating gocmd2 features and patterns.

## Timer Application

A complete example from the repository showing timer functionality with shared state and dynamic prompts.

### Full Source Code

```go
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/Necromancerlabs/gocmd2/pkg/shell"
    "github.com/Necromancerlabs/gocmd2/pkg/shellapi"
    "github.com/spf13/cobra"
)

// TimerModule provides time-tracking commands
type TimerModule struct {
    shell shellapi.ShellAPI
}

func NewTimerModule() *TimerModule {
    return &TimerModule{}
}

func (m *TimerModule) Name() string {
    return "timer"
}

func (m *TimerModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
    m.shell.SetState("start_time", time.Now())
}

func (m *TimerModule) GetCommands() []*cobra.Command {
    timeCmd := &cobra.Command{
        Use:   "time",
        Short: "Show elapsed time since shell started",
        Run: func(cmd *cobra.Command, args []string) {
            startTimeValue, ok := m.shell.GetState("start_time")
            if !ok {
                fmt.Println("Start time not found")
                return
            }

            startTime := startTimeValue.(time.Time)
            elapsed := time.Since(startTime)
            fmt.Printf("Shell running for %s\n", elapsed.Round(time.Second))

            // Update prompt with elapsed time
            m.shell.SetPrompt(fmt.Sprintf("(%s) ", elapsed.Round(time.Second)))
        },
    }

    resetCmd := &cobra.Command{
        Use:   "reset",
        Short: "Reset the timer",
        Run: func(cmd *cobra.Command, args []string) {
            m.shell.SetState("start_time", time.Now())
            m.shell.SetPrompt("> ")
            fmt.Println("Timer reset!")
        },
    }

    return []*cobra.Command{timeCmd, resetCmd}
}

func main() {
    sh, err := shell.NewShell("timer-demo", "Welcome to Timer Demo!")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
    defer sh.Close()

    // Register exit handler
    sh.OnExit(func() {
        fmt.Println("Goodbye!")
    })

    // Register timer module
    sh.RegisterModule(NewTimerModule())

    // Run shell
    sh.Run()
}
```

### Running the Example

```bash
cd gocmd2
go run examples/simple/main.go
```

### Sample Session

```
Welcome to Timer Demo!
> help
Available Commands (by module):

[core]
  exit      Exit the shell
  modules   List all modules
  enable    Enable a module
  disable   Disable a module
  help      Display this help message

[timer]
  time      Show elapsed time since shell started
  reset     Reset the timer

> time
Shell running for 5s
(5s) > time
Shell running for 10s
(10s) > reset
Timer reset!
> exit
Goodbye!
```



## Calculator Module

A module demonstrating stateful operations.

```go
package calculator

import (
    "fmt"
    "strconv"

    "github.com/Necromancerlabs/gocmd2/pkg/shellapi"
    "github.com/spf13/cobra"
)

type CalculatorModule struct {
    shell shellapi.ShellAPI
}

func New() *CalculatorModule {
    return &CalculatorModule{}
}

func (m *CalculatorModule) Name() string {
    return "calc"
}

func (m *CalculatorModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
    m.shell.SetState("calc_result", 0.0)
}

func (m *CalculatorModule) GetCommands() []*cobra.Command {
    addCmd := &cobra.Command{
        Use:   "add [number]",
        Short: "Add to accumulator",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            num, err := strconv.ParseFloat(args[0], 64)
            if err != nil {
                fmt.Println("Invalid number")
                return
            }
            result, _ := m.shell.GetState("calc_result")
            newResult := result.(float64) + num
            m.shell.SetState("calc_result", newResult)
            fmt.Printf("Result: %.2f\n", newResult)
        },
    }

    subCmd := &cobra.Command{
        Use:   "sub [number]",
        Short: "Subtract from accumulator",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            num, err := strconv.ParseFloat(args[0], 64)
            if err != nil {
                fmt.Println("Invalid number")
                return
            }
            result, _ := m.shell.GetState("calc_result")
            newResult := result.(float64) - num
            m.shell.SetState("calc_result", newResult)
            fmt.Printf("Result: %.2f\n", newResult)
        },
    }

    clearCmd := &cobra.Command{
        Use:   "clear",
        Short: "Clear accumulator",
        Run: func(cmd *cobra.Command, args []string) {
            m.shell.SetState("calc_result", 0.0)
            fmt.Println("Cleared")
        },
    }

    resultCmd := &cobra.Command{
        Use:   "result",
        Short: "Show current result",
        Run: func(cmd *cobra.Command, args []string) {
            result, _ := m.shell.GetState("calc_result")
            fmt.Printf("Result: %.2f\n", result.(float64))
        },
    }

    return []*cobra.Command{addCmd, subCmd, clearCmd, resultCmd}
}
```



## File Operations Module

A module demonstrating file system interactions.

```go
package files

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/Necromancerlabs/gocmd2/pkg/shellapi"
    "github.com/spf13/cobra"
)

type FilesModule struct {
    shell shellapi.ShellAPI
}

func New() *FilesModule {
    return &FilesModule{}
}

func (m *FilesModule) Name() string {
    return "files"
}

func (m *FilesModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
    cwd, _ := os.Getwd()
    m.shell.SetState("cwd", cwd)
}

func (m *FilesModule) GetCommands() []*cobra.Command {
    pwdCmd := &cobra.Command{
        Use:   "pwd",
        Short: "Print working directory",
        Run: func(cmd *cobra.Command, args []string) {
            cwd, _ := m.shell.GetState("cwd")
            fmt.Println(cwd.(string))
        },
    }

    cdCmd := &cobra.Command{
        Use:   "cd [path]",
        Short: "Change directory",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            cwd, _ := m.shell.GetState("cwd")
            newPath := filepath.Join(cwd.(string), args[0])

            if info, err := os.Stat(newPath); err != nil || !info.IsDir() {
                fmt.Printf("Not a directory: %s\n", args[0])
                return
            }

            absPath, _ := filepath.Abs(newPath)
            m.shell.SetState("cwd", absPath)
            m.shell.SetPrompt(fmt.Sprintf("%s> ", filepath.Base(absPath)))
        },
    }

    lsCmd := &cobra.Command{
        Use:   "ls",
        Short: "List directory contents",
        Run: func(cmd *cobra.Command, args []string) {
            cwd, _ := m.shell.GetState("cwd")
            entries, err := os.ReadDir(cwd.(string))
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
            }

            for _, entry := range entries {
                if entry.IsDir() {
                    fmt.Printf("%s/\n", entry.Name())
                } else {
                    fmt.Println(entry.Name())
                }
            }
        },
    }

    return []*cobra.Command{pwdCmd, cdCmd, lsCmd}
}
```



## Multi-Module Application

Combining multiple modules in one shell.

```go
package main

import (
    "fmt"
    "os"

    "github.com/Necromancerlabs/gocmd2/pkg/shell"
    "myapp/modules/calculator"
    "myapp/modules/files"
    "myapp/modules/timer"
)

func main() {
    sh, err := shell.NewShell("myapp", `
Welcome to MyApp!
Type 'help' for available commands.
`)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
    defer sh.Close()

    // Register multiple modules
    sh.RegisterModule(timer.New())
    sh.RegisterModule(calculator.New())
    sh.RegisterModule(files.New())

    // Cleanup on exit
    sh.OnExit(func() {
        fmt.Println("Saving state...")
    })

    sh.Run()
}
```



## Command with Flags

Using Cobra's flag system.

```go
func (m *MyModule) GetCommands() []*cobra.Command {
    var verbose bool
    var count int

    cmd := &cobra.Command{
        Use:   "process [file]",
        Short: "Process a file",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            if verbose {
                fmt.Printf("Processing %s with count %d\n", args[0], count)
            }
            // Process logic here
        },
    }

    cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
    cmd.Flags().IntVarP(&count, "count", "c", 1, "Number of iterations")

    return []*cobra.Command{cmd}
}
```

**Usage:**
```
> process myfile.txt -v -c 5
Processing myfile.txt with count 5
```



## Testing Commands

Execute commands programmatically for testing.

```go
func TestMyCommand(t *testing.T) {
    sh, _ := shell.NewShell("test", "")
    defer sh.Close()

    sh.RegisterModule(mymodule.New())

    // Capture output (redirect stdout if needed)
    sh.ExecuteCommand("mycommand arg1 arg2")

    // Check state
    value, ok := sh.GetState("expected_key")
    if !ok || value != expectedValue {
        t.Errorf("Command did not set expected state")
    }
}
```
