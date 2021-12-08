package config

import (
	"os"
	"reflect"
	"testing"
)

// TestGetEnvVariable - Testing the getting of environment variables previously declared
func TestGetEnvVariable(t *testing.T) {
	os.Setenv("MODE", "TEST")
	tests := []struct {
		key   string
		name  string
		value string
	}{
		{key: "MODE", name: "Getting valid env var", value: "TEST"},
		{key: "INVALID", name: "Getting INvalid env var", value: ""},
	}

	for _, test := range tests {
		got, _ := GetEnvVariable(test.key)

		if !reflect.DeepEqual(test.value, got) {
			t.Fatalf("%s: Expect value %v but got %v", test.name, test.value, got)
		}
	}
}
