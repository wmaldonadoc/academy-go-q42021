package datastore

import (
	"os"
	"reflect"
	"testing"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
)

func TestNewCSV(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		errCode  int
	}{
		{name: "Reading a valid CSV file", filePath: "./data-test.csv", errCode: 0},
		{name: "Reading an invalid CSV file", filePath: "./data-test1.csv", errCode: constants.DefaultExceptionCode},
	}

	for _, test := range tests {
		var osvfs = os.Open

		got, err := NewCSV(test.filePath, osvfs)
		if err != nil {
			if !reflect.DeepEqual(err.Code, test.errCode) {
				t.Error(err)
				t.Fatalf("%s: The error code should be %d but got %d", test.name, test.errCode, err.Code)
			}
		}
		if err == nil && len(got) == 0 {
			t.Fatalf("%s: The slice of pokemons should not be empty but got %v", test.name, got)
		}
	}

}
