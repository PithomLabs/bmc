package report

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// WriteJSON marshals the Report struct to indented, deterministic JSON format and writes to a file.
func WriteJSON(report *Report, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}

	// Append a trailing newline for compliance with standard file formatting
	data = append(data, '\n')

	return os.WriteFile(path, data, 0644)
}
