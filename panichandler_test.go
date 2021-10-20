package panichandler

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestPanicHandler(t *testing.T) {
	handler := &random{}
	clock := &clock{}
	handlers := []Handler{
		Handler{"/random", handler},
		Handler{"/clock", clock},
	}
	ph := NewPanicHandler(handlers...)
	http.ListenAndServe(":8081", ph)
}

type random struct{}

func (h *random) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rdm := rand.Int()
	if rdm%2 == 0 {
		panic("Oops! You've encountered an internal panic for random service")
	}
	fmt.Fprintln(w, rdm)
}

type clock struct{}

func (h *clock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if rand.Int()%2 == 0 {
		panic("Oops! You've encountered an internal panic for clock service")
	}
	fmt.Fprintln(w, time.Now())
}
