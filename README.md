# MCPVault

MCPVault (`mcpv`) is a CLI tool for managing MCP server configurations in a key-value store approach. It allows developers to add, remove, list, and view configurations for their MCP servers.

## Features

- Store and manage JSON configurations
- Hybrid storage model with index and individual files
- Simple CLI interface with intuitive commands
- Atomic file operations for reliability

## Installation

### Using Homebrew

```bash
# Coming soon
brew tap mcpvault/mcpvault
brew install mcpvault
```

### Manual Installation

1. Clone the repository
2. Build the binary
   ```bash
   go build -o mcpv main.go
   ```
3. Move the binary to your PATH
   ```bash
   sudo mv mcpv /usr/local/bin/
   ```

## Usage

### Add a configuration

```bash
# From a JSON file
mcpv add config.json --name weather --description "Weather service configuration"

# From a JSON string
mcpv add '{"api_key":"abc123","endpoint":"https://api.weather.com"}' --name weather
```

### List configurations

```bash
# List all configurations
mcpv list

# List with detailed information
mcpv list --verbose
```

### Display a configuration

```bash
# Display in JSON format (default)
mcpv cat weather

# Display in other formats (coming soon)
mcpv cat weather --format=yaml
```

### Remove a configuration

```bash
# Remove with confirmation prompt
mcpv remove weather

# Force remove without confirmation
mcpv remove weather --force
```

## Storage

MCPVault uses a hybrid storage model:

- Index file (`~/.mcpvault/index.json`): Contains metadata about all configurations
- Individual config files: Stored in `~/.mcpvault/configs/` directory with `.json` extension

## Development

### Requirements

- Go 1.16 or later
- [Cobra](https://github.com/spf13/cobra) CLI library

### Building from source

```bash
go mod download
go build -o mcpv main.go
```

## License

MIT 