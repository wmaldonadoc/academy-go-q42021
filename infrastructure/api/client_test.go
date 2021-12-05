package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
)

type MockClient struct {
	GetFunc func(url string) (*http.Response, error)
}

var getDoFunc func(url string) (*http.Response, error)

type GetDoFunc func(url string) (*http.Response, error)

func (m *MockClient) Get(url string) (*http.Response, error) {
	return getDoFunc(url)
}

// TestGet - Test the wrapper of HTTP handler.
// Should return a success HTTPStatus when it's a valid request and a specific error when somethin goes wrong
func TestGet(t *testing.T) {
	tests := []struct {
		URL        string
		HTTPStatus int
		name       string
		Request    GetDoFunc
		ErrorCode  int
	}{
		{
			URL:        "https://jsonplaceholder.typicode.com/todos/1",
			HTTPStatus: http.StatusOK,
			name:       "Making a valid GET request",
			ErrorCode:  0,
		},
		{
			URL:        "http://fakeAPI.com",
			HTTPStatus: http.StatusOK,
			name:       "Making an invalid GET request",
			ErrorCode:  constants.ThirdPartAPIExceptionCode,
		},
	}

	for _, test := range tests {
		getDoFunc = func(url string) (*http.Response, error) {
			json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
			// create a new reader with that JSON
			r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}
		client := NewApiClient(&MockClient{})

		got, err := client.Get(test.URL)

		if err != nil {
			if !reflect.DeepEqual(err.Code, test.ErrorCode) {
				t.Fatalf("%s: Expect value %d but got %d", test.name, test.ErrorCode, err.Code)
			}
		}

		if !reflect.DeepEqual(got.HTTPStatus, test.HTTPStatus) {
			t.Fatalf("%s: Expect value %d but got %d", test.name, test.HTTPStatus, got.HTTPStatus)
		}

	}
}
