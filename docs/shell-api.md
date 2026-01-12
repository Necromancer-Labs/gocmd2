# Shell API

The Shell API provides methods for modules to interact with the shell. It's passed to modules during initialization.

## ShellAPI Interface

```go
type ShellAPI interface {
    // Module Management
    EnableModule(moduleName string) error
    DisableModule(moduleName string) error
    IsModuleEnabled(moduleName string) bool
    GetModules() []string
    GetEnabledModules() []string
    GetRootCmd() *cobra.Command
    GetModuleCommands() map[string][]*cobra.Command

    // State Management
    SetState(key string, value interface{})
    GetState(key string) (interface{}, bool)

    // UI Methods
    SetPrompt(prompt string)
    GetPrompt() string
    PrintAlert(message string)
}
```



## Module Management Methods

### EnableModule

Enable a previously disabled module.

```go
func (s *Shell) EnableModule(moduleName string) error
```

**Parameters:**
- `moduleName` - Name of the module to enable

**Returns:**
- `error` - Returns error if module not found

**Example:**
```go
err := m.shell.EnableModule("timer")
if err != nil {
    fmt.Printf("Failed to enable: %v\n", err)
}
```



### DisableModule

Disable a module, removing its commands from the shell.

```go
func (s *Shell) DisableModule(moduleName string) error
```

**Parameters:**
- `moduleName` - Name of the module to disable

**Returns:**
- `error` - Returns error if module not found or is the core module

**Note:** The `core` module cannot be disabled.

**Example:**
```go
err := m.shell.DisableModule("timer")
```



### IsModuleEnabled

Check if a module is currently enabled.

```go
func (s *Shell) IsModuleEnabled(moduleName string) bool
```

**Parameters:**
- `moduleName` - Name of the module to check

**Returns:**
- `bool` - `true` if enabled, `false` otherwise

**Example:**
```go
if m.shell.IsModuleEnabled("timer") {
    fmt.Println("Timer module is active")
}
```



### GetModules

Get a list of all registered module names.

```go
func (s *Shell) GetModules() []string
```

**Returns:**
- `[]string` - Slice of module names

**Example:**
```go
modules := m.shell.GetModules()
for _, name := range modules {
    fmt.Println(name)
}
```



### GetEnabledModules

Get a list of currently enabled module names.

```go
func (s *Shell) GetEnabledModules() []string
```

**Returns:**
- `[]string` - Slice of enabled module names



### GetRootCmd

Get the root Cobra command.

```go
func (s *Shell) GetRootCmd() *cobra.Command
```

**Returns:**
- `*cobra.Command` - The root command

**Use Cases:**
- Adding subcommands dynamically
- Accessing command tree



### GetModuleCommands

Get a map of module names to their commands.

```go
func (s *Shell) GetModuleCommands() map[string][]*cobra.Command
```

**Returns:**
- `map[string][]*cobra.Command` - Map of module names to command slices



## State Management Methods

The state system provides thread-safe shared storage between modules.

### SetState

Store a value in the shared state.

```go
func (s *Shell) SetState(key string, value interface{})
```

**Parameters:**
- `key` - String identifier for the value
- `value` - Any value to store

**Example:**
```go
m.shell.SetState("counter", 0)
m.shell.SetState("config", map[string]string{"mode": "debug"})
m.shell.SetState("start_time", time.Now())
```



### GetState

Retrieve a value from the shared state.

```go
func (s *Shell) GetState(key string) (interface{}, bool)
```

**Parameters:**
- `key` - String identifier for the value

**Returns:**
- `interface{}` - The stored value
- `bool` - `true` if key exists, `false` otherwise

**Example:**
```go
value, ok := m.shell.GetState("counter")
if ok {
    counter := value.(int) // Type assertion
    fmt.Printf("Counter: %d\n", counter)
}
```

**Note:** Type assertions are required when retrieving values.



## UI Methods

### SetPrompt

Change the shell prompt.

```go
func (s *Shell) SetPrompt(prompt string)
```

**Parameters:**
- `prompt` - New prompt string

**Example:**
```go
m.shell.SetPrompt("myshell> ")
m.shell.SetPrompt("[debug] > ")
m.shell.SetPrompt("(5s elapsed) > ")
```



### GetPrompt

Get the current prompt.

```go
func (s *Shell) GetPrompt() string
```

**Returns:**
- `string` - Current prompt



### PrintAlert

Print a message to the shell output.

```go
func (s *Shell) PrintAlert(message string)
```

**Parameters:**
- `message` - Message to display

**Example:**
```go
m.shell.PrintAlert("Operation completed successfully")
```



## Shell Methods

These methods are called on the Shell instance directly, not through ShellAPI.

### NewShell

Create a new shell instance.

```go
func NewShell(name string, banner string) (*Shell, error)
```

**Parameters:**
- `name` - Root command name
- `banner` - Welcome message

**Returns:**
- `*Shell` - New shell instance
- `error` - Error if initialization fails



### RegisterModule

Register a command module.

```go
func (s *Shell) RegisterModule(module module.CommandModule)
```

**Parameters:**
- `module` - Module implementing `CommandModule` interface



### Run

Start the interactive shell loop.

```go
func (s *Shell) Run()
```



### Close

Close the shell and release resources.

```go
func (s *Shell) Close()
```



### ExecuteCommand

Execute a command programmatically.

```go
func (s *Shell) ExecuteCommand(command string)
```

**Parameters:**
- `command` - Command string to execute



### OnExit

Register a cleanup function.

```go
func (s *Shell) OnExit(fn func())
```

**Parameters:**
- `fn` - Function to call on exit



### SetHistoryFile

Change the history file location.

```go
func (s *Shell) SetHistoryFile(path string)
```

**Parameters:**
- `path` - Path to history file



## Thread Safety

The state management system uses `sync.RWMutex` for thread-safe access:
- Multiple goroutines can read state simultaneously
- Write operations are exclusive
- Safe for concurrent module access
