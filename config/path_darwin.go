// +build darwin

package config

import (
	"os"
	"path"
)

func ConfigPath() string {
		home, _ := os.UserHomeDir()
		cfgpath := path.Join(home, "Library", "Preferences")
	return cfgpath
}