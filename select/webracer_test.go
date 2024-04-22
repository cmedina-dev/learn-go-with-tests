package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares the speed of two given urls", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 1*time.Second)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("times out if response takes longer than 5s", func(t *testing.T) {
		slowServer := makeDelayedServer(6 * time.Second)
		fastServer := makeDelayedServer(5 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		timeout := 1 * time.Second

		_, err := Racer(slowURL, fastURL, timeout)

		if err == nil {
			t.Errorf("expected an error")
		}
	})
}
