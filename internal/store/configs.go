package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigStore manages operations on individual config files
type ConfigStore struct {
	index *Index
}

// NewConfigStore creates a new config store
func NewConfigStore() (*ConfigStore, error) {
	index, err := NewIndex()
	if err != nil {
		return nil, err
	}

	return &ConfigStore{
		index: index,
	}, nil
}

// Add adds a new configuration
func (c *ConfigStore) Add(name, description string, content []byte) error {
	// Validate JSON content
	var jsonContent interface{}
	if err := json.Unmarshal(content, &jsonContent); err != nil {
		return fmt.Errorf("invalid JSON content: %w", err)
	}

	// Add to index first
	if err := c.index.Add(name, description); err != nil {
		return err
	}

	// Save config file
	configsDir, err := c.index.GetConfigsDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configsDir, name+".json")
	
	// Write to temporary file first
	tempFile := configPath + ".tmp"
	if err := os.WriteFile(tempFile, content, 0644); err != nil {
		// Rollback index change on failure
		c.index.Remove(name)
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Atomically replace the original file
	if err := os.Rename(tempFile, configPath); err != nil {
		// Rollback index change on failure
		c.index.Remove(name)
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
}

// Remove removes a configuration
func (c *ConfigStore) Remove(name string) error {
	// Check if config exists
	if !c.index.Exists(name) {
		return fmt.Errorf("configuration %s does not exist", name)
	}

	// Remove config file first
	configsDir, err := c.index.GetConfigsDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configsDir, name+".json")
	if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove config file: %w", err)
	}

	// Remove from index
	return c.index.Remove(name)
}

// Get retrieves a configuration by name
func (c *ConfigStore) Get(name string) ([]byte, error) {
	// Check if config exists
	if !c.index.Exists(name) {
		return nil, fmt.Errorf("configuration %s does not exist", name)
	}

	// Read config file
	configsDir, err := c.index.GetConfigsDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configsDir, name+".json")
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return content, nil
}

// List returns all configuration names and their metadata
func (c *ConfigStore) List() map[string]ConfigMetadata {
	return c.index.List()
}

// Update updates a configuration
func (c *ConfigStore) Update(name, description string, content []byte) error {
	// Check if config exists
	if !c.index.Exists(name) {
		return fmt.Errorf("configuration %s does not exist", name)
	}

	// Validate JSON content if provided
	if content != nil {
		var jsonContent interface{}
		if err := json.Unmarshal(content, &jsonContent); err != nil {
			return fmt.Errorf("invalid JSON content: %w", err)
		}
	}

	// Update index
	if err := c.index.Update(name, description); err != nil {
		return err
	}

	// If no content to update, we're done
	if content == nil {
		return nil
	}

	// Save config file
	configsDir, err := c.index.GetConfigsDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configsDir, name+".json")
	
	// Write to temporary file first
	tempFile := configPath + ".tmp"
	if err := os.WriteFile(tempFile, content, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Atomically replace the original file
	if err := os.Rename(tempFile, configPath); err != nil {
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
} 