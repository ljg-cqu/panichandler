package panichandler

import (
	"fmt"
	"log"
	"net/http"
)

type PanicHandler struct {
	router map[string]http.Handler
}

type Handler struct {
	Pattern string
	Handler http.Handler
}

func NewPanicHandler(handlers ...Handler) *PanicHandler {
	ph := &PanicHandler{
		router: make(map[string]http.Handler),
	}

	for _, v := range handlers {
		ph.router[v.Pattern] = v.Handler
	}
	return ph
}

func (p *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	if handler, ok := p.router[r.URL.Path]; !ok {
		http.Error(w, "Sorry, service unsupported currently.", http.StatusNotFound)
	} else {
		handler.ServeHTTP(w, r)
		fmt.Println("Success:", r.Method, r.URL)
	}
}
