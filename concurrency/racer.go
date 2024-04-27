package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

// Problems
// 1. The function Racer is not testable because it makes real HTTP requests.
// 2. It uses real time to measure the duration of the requests.
func Racer(a, b string, timeout time.Duration) (winner string, error error) {
	// wait on multiple channels
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
