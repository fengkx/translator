// +build !darwin
// +build !windows

package config

import (
	"os"
	"path"
)

func ConfigPath() string {
	cfgpath := os.Getenv("XDG_CONFIG_HOME ")
	if  cfgpath == "" {
		home, _ := os.UserHomeDir()
		cfgpath = path.Join(home, ".config")
	}
	return cfgpath
}