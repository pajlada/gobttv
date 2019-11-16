package gobttv

import "testing"

func onHTTPError(t *testing.T) func(statusCode int, statusMessage, errorMessage string) {
	return func(statusCode int, statusMessage, errorMessage string) {
		t.Fatalf("HTTP error %d (%s): %s", statusCode, statusMessage, errorMessage)
	}
}

func onInternalError(t *testing.T) func(err error) {
	return func(err error) {
		t.Fatal(err)
	}
}
