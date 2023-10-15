package util

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	key := "COLOR"
	value := "green"
	os.Setenv(key, value)
	defer os.Unsetenv(key)

	result := GetEnv(key, "default")
	if result != value {
		t.Errorf("getEnv(%s) = %s; want %s", key, result, value)
	}

	key = "NON_EXISTENT"
	defaultValue := "default"
	result = GetEnv(key, defaultValue)
	if result != defaultValue {
		t.Errorf("getEnv(%s) = %s; want %s", key, result, defaultValue)
	}
}
