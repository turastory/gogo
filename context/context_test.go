package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	// Some things happen while fetching data

	go func() {
		var result string

		// This mimics a long-running task
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		store := SpyStore{response: "hello, world", t: t}
		server := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		if res.Body.String() != store.response {
			t.Errorf("got %q, want %q", res.Body.String(), store.response)
		}
	})

	t.Run("handle cancelled request", func(t *testing.T) {
		store := SpyStore{response: "hello, world", t: t}
		server := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		// All contexts derived from it will be cancelled
		req = req.WithContext(cancellingContext)

		res := &SpyResponseWriter{}

		server.ServeHTTP(res, req)

		if res.written {
			t.Errorf("a response should not have been written")
		}
	})
}
