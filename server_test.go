package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test hello", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()

		Greeting(resp, req)
		got := resp.Body.String()
		want := "hello"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
