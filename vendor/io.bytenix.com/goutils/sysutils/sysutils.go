package sysutils

import (
	"os"
	"path"
)

// InitAppDataDirectory initializes the application data directory in the user home
func InitAppDataDirectory(name string) (string, error) {
	appDataPath := path.Join(os.Getenv("HOME"), ".local", "share", name)

	err := os.MkdirAll(appDataPath, 0750)

	return appDataPath, err
}
