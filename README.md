# MCPVault 🗃️

MCPVault (`mcpv`) is a stupidly simple CLI tool for managing MCP server configurations.  It allows you to add, remove, list, and view configurations for your MCP servers.

## Features
 - `mcpv add <path> <name> <description> ` - Add a configuration from a file
 - `mcpv add "{}" --name JobSearch` - Add a configuration from a JSON string
 - `mcpv remove <name>` - Remove a configuration
 - `mcpv list` - List all configurations
 - `mcpv cat <name>` - View a configuration

## Installation MacOS (Local)
```bash
# Clone the repository
git clone https://github.com/andrewlayer/MCPVault.git
cd MCPVault

# Build the binary
go build -o mcpv

# Move to a directory in your PATH (optional)
sudo mv mcpv /usr/local/bin/
```

### Using Homebrew
Well I need >75 stars to get this on Homebrew. So, if you like this project, please star it (or fork it)!