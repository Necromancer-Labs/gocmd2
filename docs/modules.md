# Modules

Modules are the primary way to extend gocmd2 functionality. Each module provides a set of related commands that can be enabled or disabled at runtime.

## The CommandModule Interface

To create a module, implement the `CommandModule` interface:

```go
type CommandModule interface {
    // GetCommands returns the Cobra commands this module provides
    GetCommands() []*cobra.Command

    // Name returns the module's identifier
    Name() string

    // Initialize is called when the module is registered
    Initialize(shell shellapi.ShellAPI)
}
```


## Creating a Module

### Basic Structure

```go
package mymodule

import (
    "github.com/Necromancerlabs/gocmd2/pkg/shellapi"
    "github.com/spf13/cobra"
)

type MyModule struct {
    shell shellapi.ShellAPI
}

func New() *MyModule {
    return &MyModule{}
}

func (m *MyModule) Name() string {
    return "mymodule"
}

func (m *MyModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
    // Setup initial state, resources, etc.
}

func (m *MyModule) GetCommands() []*cobra.Command {
    return []*cobra.Command{
        // Your commands here
    }
}
```

### Adding Commands

Commands use the [Cobra](https://github.com/spf13/cobra) library:

```go
func (m *MyModule) GetCommands() []*cobra.Command {
    greetCmd := &cobra.Command{
        Use:   "greet [name]",
        Short: "Greet someone",
        Long:  "Display a greeting message for the specified name",
        Args:  cobra.MaximumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            name := "World"
            if len(args) > 0 {
                name = args[0]
            }
            fmt.Printf("Hello, %s!\n", name)
        },
    }

    return []*cobra.Command{greetCmd}
}
```


## Registering Modules

Register modules with the shell before calling `Run()`:

```go
sh, _ := shell.NewShell("myshell", "Welcome!")

// Register custom modules
sh.RegisterModule(mymodule.New())
sh.RegisterModule(anothermodule.New())

sh.Run()
```


## Module Lifecycle

1. **Registration**: `RegisterModule()` is called
2. **Initialization**: `Initialize(shell)` is called with shell API access
3. **Commands Added**: `GetCommands()` returns commands to add
4. **Runtime**: Module can be enabled/disabled by users
5. **Cleanup**: Exit handlers run when shell closes


## Using Shared State

Modules can share data through the shell's state system:

```go
func (m *MyModule) Initialize(s shellapi.ShellAPI) {
    m.shell = s
    // Set initial state
    m.shell.SetState("counter", 0)
}

func (m *MyModule) GetCommands() []*cobra.Command {
    incrementCmd := &cobra.Command{
        Use:   "increment",
        Short: "Increment the counter",
        Run: func(cmd *cobra.Command, args []string) {
            value, _ := m.shell.GetState("counter")
            counter := value.(int)
            counter++
            m.shell.SetState("counter", counter)
            fmt.Printf("Counter: %d\n", counter)
        },
    }

    return []*cobra.Command{incrementCmd}
}
```


## Dynamic UI Updates

Modules can modify the shell prompt:

```go
func (m *MyModule) GetCommands() []*cobra.Command {
    statusCmd := &cobra.Command{
        Use:   "status",
        Short: "Show status in prompt",
        Run: func(cmd *cobra.Command, args []string) {
            m.shell.SetPrompt("[active] > ")
        },
    }

    resetCmd := &cobra.Command{
        Use:   "reset-prompt",
        Short: "Reset prompt to default",
        Run: func(cmd *cobra.Command, args []string) {
            m.shell.SetPrompt("> ")
        },
    }

    return []*cobra.Command{statusCmd, resetCmd}
}
```


## Module Enable/Disable

Users control modules at runtime:

```
> modules              # List all modules
> disable mymodule     # Disable a module (removes its commands)
> enable mymodule      # Re-enable a module (restores commands)
```

The core module cannot be disabled.


## Best Practices

1. **Unique Names**: Use descriptive, unique module names
2. **Initialize State Early**: Set up state in `Initialize()`
3. **Group Related Commands**: Keep related functionality in one module
4. **Document Commands**: Use `Short` and `Long` descriptions
5. **Handle Errors Gracefully**: Don't crash the shell on errors
6. **Clean Up Resources**: Use exit handlers for cleanup


## Example: Complete Module

```go
package timer

import (
    "fmt"
    "time"

    "github.com/Necromancerlabs/gocmd2/pkg/shellapi"
    "github.com/spf13/cobra"
)

type TimerModule struct {
    shell shellapi.ShellAPI
}

func New() *TimerModule {
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
        Short: "Show elapsed time",
        Run: func(cmd *cobra.Command, args []string) {
            start, _ := m.shell.GetState("start_time")
            elapsed := time.Since(start.(time.Time))
            fmt.Printf("Elapsed: %s\n", elapsed.Round(time.Second))
            m.shell.SetPrompt(fmt.Sprintf("(%s) > ", elapsed.Round(time.Second)))
        },
    }

    resetCmd := &cobra.Command{
        Use:   "reset",
        Short: "Reset the timer",
        Run: func(cmd *cobra.Command, args []string) {
            m.shell.SetState("start_time", time.Now())
            m.shell.SetPrompt("> ")
            fmt.Println("Timer reset")
        },
    }

    return []*cobra.Command{timeCmd, resetCmd}
}
```
