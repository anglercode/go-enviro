package enviro

import (
	"os"
	"strconv"
	"strings"
)

// Get returns the value of an environment variable or the provided default when non-existant.
func Get(name string, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}

// GetInt returns the value of an environment variable as an integer value.
func GetInt(name string, defaultValue int) int {
	value := Get(name, "")
	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	return defaultValue
}

// GetBool returns the value of an environment variable as a boolean value.
func GetBool(name string, defaultValue bool) bool {
	value := Get(name, "")
	if val, err := strconv.ParseBool(value); err == nil {
		return val
	}
	return defaultValue
}

// GetSlice returns the value of an environment variable as a string slice based on the provided seperator.
func GetSlice(name string, sep string, defaultValue []string) []string {
	value := Get(name, "")

	if value == "" {
		return defaultValue
	}
	return strings.Split(value, sep)
}
