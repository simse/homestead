package config

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	os.MkdirAll(ConfigLocation(), os.ModePerm)
}

// ConfigLocation finds user home directory and returns path
func ConfigLocation() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dirname, ".homestead")
}

// ConfigFileLocation returns path to config file
func ConfigFileLocation() string {
	return filepath.Join(ConfigLocation(), "homestead.db")
}
