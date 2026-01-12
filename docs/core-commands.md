---
layout: default
title: Core Commands
nav_order: 5
---

# Core Commands

The core module is automatically registered with every gocmd2 shell. It provides essential commands for shell operation and module management.

The core module cannot be disabled.

---

## help

Display available commands grouped by module.

```
> help
```

**Output:**
- Lists all commands organized by their parent module
- Shows command descriptions
- Indicates which modules are disabled

**Example Output:**
```
Available Commands (by module):

[core]
  exit        Exit the shell
  modules     List all modules
  enable      Enable a module
  disable     Disable a module
  help        Display this help message

[timer]
  time        Show elapsed time
  reset       Reset the timer

[Disabled Modules]
  mymodule (use 'enable mymodule' to activate)
```

---

## exit

Exit the shell gracefully.

```
> exit
```

**Behavior:**
- Triggers all registered exit handlers
- Closes readline instance
- Returns control to the calling program

**Note:** You can also exit with `Ctrl+D` (EOF).

---

## modules

List all registered modules and their status.

```
> modules
```

**Output:**
- Module name
- Status (enabled/disabled)

**Example Output:**
```
Available modules:
  core     [enabled]
  timer    [enabled]
  utils    [disabled]
```

---

## enable

Enable a disabled module.

```
> enable <module_name>
```

**Parameters:**
- `module_name` - Name of the module to enable

**Example:**
```
> enable timer
Module 'timer' enabled
```

**Behavior:**
- Restores the module's commands to the shell
- Commands become available immediately
- Tab completion updates

**Errors:**
- "Module not found" if module doesn't exist

---

## disable

Disable an enabled module.

```
> disable <module_name>
```

**Parameters:**
- `module_name` - Name of the module to disable

**Example:**
```
> disable timer
Module 'timer' disabled
```

**Behavior:**
- Removes the module's commands from the shell
- Commands are no longer available
- Tab completion updates
- Module state is preserved

**Restrictions:**
- Cannot disable the `core` module

**Errors:**
- "Cannot disable core module"
- "Module not found" if module doesn't exist

---

## Command Parsing

gocmd2 uses [google/shlex](https://github.com/google/shlex) for shell-like argument parsing:

### Supported Syntax

| Syntax | Example | Result |
|--------|---------|--------|
| Simple args | `cmd arg1 arg2` | `["arg1", "arg2"]` |
| Double quotes | `cmd "hello world"` | `["hello world"]` |
| Single quotes | `cmd 'hello world'` | `["hello world"]` |
| Escaped quotes | `cmd "say \"hi\""` | `["say \"hi\""]` |
| Mixed | `cmd arg1 "arg 2"` | `["arg1", "arg 2"]` |

### Examples

```
> greet "John Doe"
Hello, John Doe!

> echo 'single quoted'
single quoted

> cmd "nested \"quotes\""
nested "quotes"
```

---

## Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Tab` | Auto-complete command |
| `Ctrl+C` | Cancel current input |
| `Ctrl+D` | Exit shell (EOF) |
| `Up Arrow` | Previous command in history |
| `Down Arrow` | Next command in history |
| `Ctrl+R` | Reverse search history |
| `Ctrl+A` | Move cursor to start |
| `Ctrl+E` | Move cursor to end |
| `Ctrl+W` | Delete word backward |
| `Ctrl+U` | Clear line |

---

## Command History

Commands are automatically saved to history:

- **Default location:** `/tmp/readline.tmp`
- **Persistent:** History survives shell restarts
- **Navigation:** Use up/down arrows
- **Search:** Use `Ctrl+R` for reverse search

### Changing History Location

```go
sh.SetHistoryFile("/home/user/.myshell_history")
```
