package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// ErrInvalidConfig is returned when config is invalid
var ErrInvalidConfig = errors.New("invalid config")

// Config represents configuration for the application
type Config struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadConfig loads configuration from a file
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, ErrInvalidConfig
	}
	return &config, nil
}

// SaveConfig saves configuration to a file
func SaveConfig(filename string, config *Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// GetUUID generates a random UUID
func GetUUID() string {
	return uuid.New().String()
}

// GetWorkingDirectory returns the working directory
func GetWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

// IsDir checks if a path is a directory
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// IsFile checks if a path is a file
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

// GetFileSize returns the size of a file in bytes
func GetFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// DownloadFile downloads a file from a URL
func DownloadFile(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fcopy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// GetBaseName returns the base name of a path
func GetBaseName(path string) string {
	return filepath.Base(path)
}

// GetDirName returns the directory name of a path
func GetDirName(path string) string {
	return filepath.Dir(path)
}

// HasSuffix checks if a string has a suffix
func HasSuffix(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}