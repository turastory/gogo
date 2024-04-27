package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("select the fastest url", func(t *testing.T) {
		slowServer := makeServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl, 10*time.Second)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns timeout error when there's no response within 10s", func(t *testing.T) {
		a := makeServer(25 * time.Millisecond)
		defer a.Close()

		aUrl := a.URL

		_, err := Racer(aUrl, aUrl, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
