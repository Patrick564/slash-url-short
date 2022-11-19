package utils

import (
	"reflect"
	"testing"
)

func AssertStatusCode(t testing.TB, want, got int) {
	t.Helper()

	if want != got {
		t.Fatalf("got %d code, but want %d code", got, want)
	}
}

func AssertResponseBody(t testing.TB, want []byte, got string) {
	t.Helper()

	if !reflect.DeepEqual(string(want), got) {
		t.Fatalf("got %+v body, but want %+v body", got, string(want))
	}
}
