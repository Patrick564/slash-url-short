package models

import (
	"context"
	"os"
	"reflect"
	"testing"
)

var (
	ctx  context.Context
	conn UrlModel
	err  error
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	conn, err = OpenDatabaseConn(ctx, "")

	code := m.Run()

	conn.DB.FlushAll(ctx)
	os.Exit(code)
}

func TestModelsAdd(t *testing.T) {
	cases := []struct {
		name     string
		sid      string
		original string
		want     error
	}{
		{name: "insert url:AFCV1", sid: "AFCV1", original: "https://example1.com", want: nil},
		{name: "insert url:AFCV2", sid: "AFCV2", original: "https://example2.com", want: nil},
		{name: "insert url:AFCV3", sid: "AFCV3", original: "https://example3.com", want: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := conn.Add(c.sid, c.original)
			if err != c.want {
				t.Fatalf("error at insert to database: %+v", err)
			}
		})
	}
}

func TestModelsSearchById(t *testing.T) {
	cases := []struct {
		name string
		id   string
		want string
	}{
		{name: "get url by hash", id: "AFCV1", want: "https://example1.com"},
		{name: "get url by hash", id: "AFCV2", want: "https://example2.com"},
		{name: "get url by hash", id: "AFCV3", want: "https://example3.com"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s, err := conn.GoTo(c.id)
			if err != nil {
				t.Fatalf("error at query to database: %+v", err)
			}

			if c.want != s {
				t.Fatalf("want %s, but got %s", c.want, s)
			}
		})
	}
}

func TestModelsAll(t *testing.T) {
	want := []string{"url:AFCV2", "url:AFCV3", "url:AFCV1"}

	t.Run("query all values", func(t *testing.T) {
		k, err := conn.All()
		if err != nil {
			t.Fatalf("error at query all: %+v", err)
		}

		if !reflect.DeepEqual(want, k) {
			t.Fatalf("want %s, but got %s", want, k)
		}
	})
}
