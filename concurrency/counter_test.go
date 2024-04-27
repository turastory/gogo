package concurrency

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("naive counter", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(counter, t, 3)
	})

	t.Run("concurrent counter", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// WaitGroup is a tool to synchronize multiple goroutines together
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(counter, t, wantedCount)
	})
}

func assertCounter(counter *Counter, t *testing.T, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), 3)
	}
}
