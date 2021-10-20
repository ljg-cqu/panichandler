package panichandler

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestPanicHandlerFunc(t *testing.T) {
	http.HandleFunc("/random", PanicHandlerFunc(randomFunc))
	http.HandleFunc("/clock", PanicHandlerFunc(clockFunc))
	http.ListenAndServe(":8081", nil)
}

func randomFunc(w http.ResponseWriter, r *http.Request) {
	rdm := rand.Int()
	if rdm%2 == 0 {
		panic("Oops! You've encountered an internal panic for random service")
	}
	fmt.Fprintln(w, rdm)
}

func clockFunc(w http.ResponseWriter, r *http.Request) {
	if rand.Int()%2 == 0 {
		panic("Oops! You've encountered an internal panic for clock service")
	}
	fmt.Fprintln(w, time.Now())
}
