// +build windows

package config

import (
	"os"
	"path"
)

func ConfigPath() string {
	cfgpath := os.Getenv("APPDATA")
	if cfgpath == "" {
		home, _ := os.UserHomeDir()
		cfgpath = path.Join(home, "AppData", "Roaming")
	}
	return cfgpath
}
