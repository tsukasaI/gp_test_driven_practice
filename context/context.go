package context

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response string
	canceled bool
	t        *testing.T
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}
