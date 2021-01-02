package enviro

import (
	"os"
	"strconv"
	"strings"
)

// Get returns the value of an environment variable or
// the provided default when non-existent as a string.
func Get(name string, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}

// GetInt returns the value of an environment variable as
// an integer value or the provided default integer.
func GetInt(name string, defaultValue int) int {
	value := Get(name, "")
	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	return defaultValue
}

// GetBool returns the value of an environment variable as
// a boolean value or the provided default boolean.
func GetBool(name string, defaultValue bool) bool {
	value := Get(name, "")
	if val, err := strconv.ParseBool(value); err == nil {
		return val
	}
	return defaultValue
}

// GetSlice returns the value of an environment variable as
// a string slice based on the provided separator. It also takes
// a string slice to use as a fallback default.
func GetSlice(name string, sep string, defaultValue []string) []string {
	value := Get(name, "")

	if value == "" {
		return defaultValue
	}
	return strings.Split(value, sep)
}
