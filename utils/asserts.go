package utils

import (
	"errors"
	"reflect"
	"testing"
)

func AssertError(t testing.TB, want, got error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Fatalf("got %+v error, but want %+v error", got, want)
	}
}

func AssertResponseBody(t testing.TB, want []byte, got string) {
	t.Helper()

	if !reflect.DeepEqual(string(want), got) {
		t.Fatalf("got %+v body, but want %+v body", got, string(want))
	}
}

func AssertStatusCode(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("got %d code, but want %d code", got, want)
	}
}
