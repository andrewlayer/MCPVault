package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mcpvault/mcpvault/internal/store"
)

// Manager handles configuration operations
type Manager struct {
	store *store.ConfigStore
}

// NewManager creates a new configuration manager
func NewManager() (*Manager, error) {
	store, err := store.NewConfigStore()
	if err != nil {
		return nil, err
	}

	return &Manager{
		store: store,
	}, nil
}

// AddFromFile adds a configuration from a file
func (m *Manager) AddFromFile(name, description, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return m.store.Add(name, description, content)
}

// AddFromJSON adds a configuration from a JSON string
func (m *Manager) AddFromJSON(name, description, jsonStr string) error {
	// Validate JSON
	var jsonContent interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonContent); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	// Format JSON for consistent storage
	formattedJSON, err := json.MarshalIndent(jsonContent, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	return m.store.Add(name, description, formattedJSON)
}

// Remove removes a configuration
func (m *Manager) Remove(name string) error {
	return m.store.Remove(name)
}

// Get retrieves a configuration by name
func (m *Manager) Get(name string) ([]byte, error) {
	return m.store.Get(name)
}

// List returns all configuration names and their metadata
func (m *Manager) List(verbose bool) ([]string, error) {
	configs := m.store.List()
	
	var result []string
	if verbose {
		for name, metadata := range configs {
			result = append(result, fmt.Sprintf("%s\n  Description: %s\n  Created: %s\n  Updated: %s",
				name, metadata.Description, metadata.Created.Format("2006-01-02 15:04:05"), metadata.Updated.Format("2006-01-02 15:04:05")))
		}
	} else {
		for name := range configs {
			result = append(result, name)
		}
	}

	return result, nil
}

// ProcessInput determines if the input is a file path or JSON string
func (m *Manager) ProcessInput(input string) ([]byte, error) {
	// Check if input is a file path
	if _, err := os.Stat(input); err == nil {
		content, err := os.ReadFile(input)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		return content, nil
	}

	// Trim input and check if it looks like JSON
	input = strings.TrimSpace(input)
	if (strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}")) ||
		(strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]")) {
		// Validate JSON
		var jsonContent interface{}
		if err := json.Unmarshal([]byte(input), &jsonContent); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}

		// Format JSON for consistent storage
		formattedJSON, err := json.MarshalIndent(jsonContent, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to format JSON: %w", err)
		}

		return formattedJSON, nil
	}

	return nil, fmt.Errorf("input is neither a valid file path nor valid JSON")
}

// FormatJSON formats JSON content with indentation for display
func (m *Manager) FormatJSON(content []byte) (string, error) {
	var data interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	return string(formattedJSON), nil
}

// PrintConfig prints a configuration to the specified writer
func (m *Manager) PrintConfig(w io.Writer, name string, format string) error {
	content, err := m.store.Get(name)
	if err != nil {
		return err
	}

	switch strings.ToLower(format) {
	case "json":
		formatted, err := m.FormatJSON(content)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\n", formatted)
	case "yaml":
		return fmt.Errorf("YAML format not yet implemented")
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	return nil
} 