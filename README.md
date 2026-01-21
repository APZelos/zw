# zw

A git worktree navigation tool with fuzzy matching and interactive selection.

## Installation

### Using Go

```bash
go install github.com/apzelos/zw/cmd/zw@latest
```

### From Source

```bash
git clone https://github.com/apzelos/zw.git
cd zw
make build
```

## Shell Integration

Add the following to your shell configuration file:

**Bash** (`~/.bashrc`):
```bash
eval "$(zw init bash)"
```

**Zsh** (`~/.zshrc`):
```bash
eval "$(zw init zsh)"
```

**Fish** (`~/.config/fish/config.fish`):
```fish
zw init fish | source
```

## Usage

After setting up shell integration, simply type:

```bash
zw <pattern>
```

This will fuzzy match against your git worktrees and navigate to the best match.

### Direct Commands

Access the binary directly using `command zw`:

```bash
command zw --help      # Show help
command zw version     # Show version
command zw init bash   # Output shell integration
```

## Development

```bash
make build    # Build binary to bin/zw
make test     # Run tests
make lint     # Run linter
make run      # Run without building
```

## License

GPL-3.0 - see [LICENSE](LICENSE) for details.
