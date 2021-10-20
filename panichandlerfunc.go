package panichandler

import (
	"log"
	"net/http"
)

func PanicHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rc := recover(); rc != nil {
				des, ok := rc.(string)
				if ok {
					http.Error(w, des, 500)
				} else {
					http.Error(w, "internal panic", 500)
				}
				log.Println("panic caught:", r.Method, r.URL)
			}
		}()
		h(w, r)
	}
}
