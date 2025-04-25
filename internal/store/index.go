package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ConfigMetadata stores metadata about a configuration
type ConfigMetadata struct {
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// IndexFile represents the structure of the index.json file
type IndexFile struct {
	Configs map[string]ConfigMetadata `json:"configs"`
}

// Index manages operations on the index file
type Index struct {
	path string
	data IndexFile
}

// NewIndex creates a new index manager
func NewIndex() (*Index, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	mcpvDir := filepath.Join(homeDir, ".mcpvault")
	if err := os.MkdirAll(mcpvDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", mcpvDir, err)
	}

	configsDir := filepath.Join(mcpvDir, "configs")
	if err := os.MkdirAll(configsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", configsDir, err)
	}

	indexPath := filepath.Join(mcpvDir, "index.json")

	index := &Index{
		path: indexPath,
		data: IndexFile{
			Configs: make(map[string]ConfigMetadata),
		},
	}

	// Load existing index file if it exists
	if _, err := os.Stat(indexPath); err == nil {
		if err := index.load(); err != nil {
			return nil, fmt.Errorf("failed to load index file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check index file: %w", err)
	}

	return index, nil
}

// load reads the index file into memory
func (i *Index) load() error {
	data, err := os.ReadFile(i.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &i.data)
}

// save writes the index data to the index file
func (i *Index) save() error {
	data, err := json.MarshalIndent(i.data, "", "  ")
	if err != nil {
		return err
	}

	// Write to temporary file first
	tempFile := i.path + ".tmp"
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return err
	}

	// Atomically replace the original file
	return os.Rename(tempFile, i.path)
}

// Add adds a new configuration to the index
func (i *Index) Add(name, description string) error {
	if _, exists := i.data.Configs[name]; exists {
		return fmt.Errorf("configuration %s already exists", name)
	}

	now := time.Now().UTC()
	i.data.Configs[name] = ConfigMetadata{
		Description: description,
		Created:     now,
		Updated:     now,
	}

	return i.save()
}

// Update updates a configuration's metadata in the index
func (i *Index) Update(name, description string) error {
	metadata, exists := i.data.Configs[name]
	if !exists {
		return fmt.Errorf("configuration %s does not exist", name)
	}

	if description != "" {
		metadata.Description = description
	}
	metadata.Updated = time.Now().UTC()
	i.data.Configs[name] = metadata

	return i.save()
}

// Remove removes a configuration from the index
func (i *Index) Remove(name string) error {
	if _, exists := i.data.Configs[name]; !exists {
		return fmt.Errorf("configuration %s does not exist", name)
	}

	delete(i.data.Configs, name)
	return i.save()
}

// List returns all configuration names and their metadata
func (i *Index) List() map[string]ConfigMetadata {
	return i.data.Configs
}

// Exists checks if a configuration exists in the index
func (i *Index) Exists(name string) bool {
	_, exists := i.data.Configs[name]
	return exists
}

// GetConfigsDir returns the path to the configs directory
func (i *Index) GetConfigsDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(homeDir, ".mcpvault", "configs"), nil
} 