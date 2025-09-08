// In cache/cache.go
package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ProjectsCache holds the mapping of project names to their paths.
type ProjectsCache map[string]string

// getCachePath returns the path to the projects.json cache file.
func getCachePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", "portal", "projects.json"), nil
}

// Read loads the project cache from the JSON file.
func Read() (ProjectsCache, error) {
	path, err := getCachePath()
	if err != nil {
		return nil, err
	}

	// If the file doesn't exist, return an empty cache.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return make(ProjectsCache), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cache ProjectsCache
	err = json.Unmarshal(data, &cache)
	return cache, err
}

// Write saves the given cache to the JSON file.
func (pc ProjectsCache) Write() error {
	path, err := getCachePath()
	if err != nil {
		return err
	}

	// Ensure the directory exists.
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(pc, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
