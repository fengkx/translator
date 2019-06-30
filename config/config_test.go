package config

import (
	"testing"
)

func TestConfigPath(t *testing.T) {
	cp := ConfigPath()
	stringNoEmpty(t, cp, "Config path")
}

func stringNoEmpty(t *testing.T, s string, name ...string) {
	if len(s) < 0 {
		if len(name) > 0 {
			t.Fatalf("%s should not empty", name[0])
		}
		t.Fatal("string should not empty")
	}
}
