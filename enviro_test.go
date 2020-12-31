package enviro

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// GetEnv tests
func TestGetExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := getRandString(12)
	os.Setenv(tmpName, tmpValue)

	value := Get(tmpName, "SHOULD_NOT_BE_THIS")
	if value != tmpValue || value == "SHOULD_NOT_BE_THIS" || value == "" {
		t.Errorf("Failed to get value for an existing environment variable. Got: %s", value)
	}

}

func TestGetNonExist(t *testing.T) {
	shouldntExist := getRandString(48)
	value := Get(shouldntExist, "KNOWN_VALUE")
	if value != "KNOWN_VALUE" {
		t.Errorf("Failed to get default value from getEnv when non-existent variable is asked for. Got: %s", value)
	}
}

// GetEnvInt tests
func TestGetIntExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "666"
	os.Setenv(tmpName, tmpValue)

	value := GetInt(tmpName, 6661)
	if value != 666 {
		t.Errorf("Unable to get value for an existing environment variable as an integer. Got: %d", value)
	}
}

func TestGetIntNonExist(t *testing.T) {
	shouldntExist := getRandString(48)
	value := GetInt(shouldntExist, 1999)
	if value != 1999 {
		t.Errorf("Did not get default value from getEnvInt when non-existent variable is asked for. Got: %d", value)
	}
}

// GetEnvBool tests

func TestGetBoolExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "false"
	os.Setenv(tmpName, tmpValue)

	value := GetBool(tmpName, true)
	if value {
		t.Errorf("Unable to get value as boolean for existing environment variable.")
	}
}

func TestGetBoolExist2(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "False"
	os.Setenv(tmpName, tmpValue)

	value := GetBool(tmpName, true)
	if value {
		t.Errorf("Unable to get value as boolean for existing environment variable.")
	}
}

func TestGetBoolNonExist(t *testing.T) {
	shouldntExist := getRandString(48)
	value := GetBool(shouldntExist, false)
	if value {
		t.Errorf("Did not get default value from getEnvInt when non-existent variable is asked for.")
	}
}

// GetEnvSlice tests

func TestGetSliceCommasExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "admin,tester,user"
	os.Setenv(tmpName, tmpValue)

	value := GetSlice(tmpName, ",", []string{"not_this", "or_this"})
	if value[0] != "admin" || value[1] != "tester" || value[2] != "user" {
		t.Errorf("Unable to get value as slice for existing environment variable. Got: %s", value)
	}
}

func TestGetSliceColonsExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "admin:tester:user"
	os.Setenv(tmpName, tmpValue)

	value := GetSlice(tmpName, ":", []string{"not_this", "or_this"})
	if value[0] != "admin" || value[1] != "tester" || value[2] != "user" {
		t.Errorf("Unable to get value as slice for existing environment variable. Got: %s", value)
	}
}
func TestGetSliceNonExist(t *testing.T) {
	shouldntExist := getRandString(48)
	value := GetSlice(shouldntExist, ",", []string{"default", "tester"})
	if value[0] != "default" || value[1] != "tester" {
		t.Errorf("Did not get default values from getEnvSlice when non-existent variable is asked for. Got: %s", value)
	}
}

// Test Helpers

func getRandString(length int) string {
	var randomSeed *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randomSeed.Intn(len(charset))]
	}
	return string(b)
}
