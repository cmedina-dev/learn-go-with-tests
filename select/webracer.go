package main

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			panic("Unhandled error exception")
		}
		close(ch)
	}()
	return ch
}
