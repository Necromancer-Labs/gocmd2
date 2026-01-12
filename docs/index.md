# gocmd2

**Interactive Shell Framework** — cmd2, but for Go.

A modular, extensible framework for building interactive command-line applications with tab completion, history, and runtime module management.

## Why gocmd2?

Python has `cmd2`. Go had nothing comparable. gocmd2 gives you the same developer experience with Go's performance and deployment simplicity.

- **Modular architecture** — organize commands into enable/disable modules
- **Built on Cobra** — leverage the most popular Go CLI framework
- **Tab completion** — out of the box via readline
- **Shared state** — thread-safe state sharing between modules

## Quick Links

- [Quick Start](quickstart.md) — Get up and running
- [Modules](modules.md) — Creating command modules
- [Shell API](shell-api.md) — Full API reference
- [Core Commands](core-commands.md) — Built-in commands
- [Examples](examples.md) — Code examples

## Dependencies

| Package | Purpose |
|---------|---------|
| [spf13/cobra](https://github.com/spf13/cobra) | Command framework |
| [chzyer/readline](https://github.com/chzyer/readline) | Interactive input |
| [google/shlex](https://github.com/google/shlex) | Argument parsing |

## License

MIT — [Necromancer Labs](https://github.com/Necromancer-Labs)
