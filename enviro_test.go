package enviro

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// Get tests
func TestGetExist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := getRandString(12)
	os.Setenv(tmpName, tmpValue)

	value := Get(tmpName, "SHOULD_NOT_BE_THIS")
	if value != string(tmpValue) || value == "SHOULD_NOT_BE_THIS" || value == "" {
		t.Errorf("Failed to get value for an existing environment variable. Got: %s", value)
	}
}

func TestGetNonExist(t *testing.T) {
	shouldNotExist := getRandString(48)
	value := Get(shouldNotExist, "KNOWN_VALUE")
	if value != string("KNOWN_VALUE") {
		t.Errorf("Failed to get default value from Get when non-existent variable is asked for. Got: %s", value)
	}
}

// GetInt tests
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
	shouldNotExist := getRandString(48)
	value := GetInt(shouldNotExist, 1999)
	if value != 1999 {
		t.Errorf("Did not get default value from GetInt when non-existent variable is asked for. Got: %d", value)
	}
}

// GetInt64 tests
func TestGetInt64Exist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "6664"
	os.Setenv(tmpName, tmpValue)

	value := GetInt64(tmpName, 6661)
	if value != 6664 {
		t.Errorf("Unable to get value for an existing environment variable as an int64. Got: %d", value)
	}
}

func TestGetInt64NonExist(t *testing.T) {
	shouldNotExist := getRandString(48)
	value := GetInt64(shouldNotExist, 1999)
	if value != 1999 {
		t.Errorf("Did not get default value from GetInt64 when non-existent variable is asked for. Got: %d", value)
	}
}

// GetBool tests
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
	shouldNotExist := getRandString(48)
	value := GetBool(shouldNotExist, false)
	if value {
		t.Errorf("Did not get default value from getEnvInt when non-existent variable was asked for.")
	}
}

// GetFloat32 tests
func TestGetFloat32Exist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "5.55"
	os.Setenv(tmpName, tmpValue)

	value := GetFloat32(tmpName, 5.55)
	if value != float32(5.55) {
		t.Errorf("Unable to get value as float32 for existing environment variable.")
	}
}

func TestGetFloat32NonExist(t *testing.T) {
	shouldNotExist := getRandString(48)
	value := GetFloat32(shouldNotExist, 5.55)
	if value != float32(5.55) {
		t.Errorf("Did not get default value from GetFloat32 when non-existent variable was asked for.")
	}
}

// Get Float64 tests
func TestGetFloat64Exist(t *testing.T) {
	tmpName := getRandString(8)
	tmpValue := "5.55"
	os.Setenv(tmpName, tmpValue)

	value := GetFloat64(tmpName, 5.55)
	if value != float64(5.55) {
		t.Errorf("Unable to get value as float32 for existing environment variable.")
	}
}

func TestGetFloat64NonExist(t *testing.T) {
	shouldNotExist := getRandString(64)
	value := GetFloat64(shouldNotExist, 5.55)
	if value != float64(5.55) {
		t.Errorf("Did not get default value from GetFloat64 when non-existent variable was asked for.")
	}
}

// GetSlice tests

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
	shouldNotExist := getRandString(48)
	value := GetSlice(shouldNotExist, ",", []string{"default", "tester"})
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
