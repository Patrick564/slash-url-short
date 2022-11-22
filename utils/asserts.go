package utils

import (
	"reflect"
	"testing"
)

func AssertError(t testing.TB, want, got string) {
	t.Helper()

	if got == want {
		t.Fatalf("got %+v error, but want %+v error", got, want)
	}
}

func AssertResponseBody(t testing.TB, want, got string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %+v body, but want %+v body", got, string(want))
	}
}

func AssertStatusCode(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("got %d code, but want %d code", got, want)
	}
}
