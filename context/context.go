package context

import (
	"context"
	"fmt"
	"net/http"
)

// This represents a store that can fetch data, potentially
// from a database or an external API.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, res)
	}
}
