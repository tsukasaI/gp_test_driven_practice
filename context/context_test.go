package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}
func TestHandler(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
		store.assertWasNotCanceled()
	})

	t.Run("tells store to cancel work if request is canceled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCanceled()
	})

}

func (s *SpyStore) assertWasCanceled() {
	s.t.Helper()
	if !s.canceled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCanceled() {
	s.t.Helper()
	if s.canceled {
		s.t.Error("store was told to cancel")
	}
}
