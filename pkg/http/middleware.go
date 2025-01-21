package http

import (
	"net/http"

	"github.com/joseluis8906/pocone/pkg/log"
)

func LogReq(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request received: %s %s", log.Info, r.Method, r.URL.Path)
}

func LogRes(w http.ResponseWriter, r *http.Request) {
	rw, ok := w.(*responseWriter)
	if ok {
		log.Printf("%s response sent: %s %s [%d] [size=%d]", log.Info, r.Method, r.URL.Path, rw.code, rw.size)
	}
}
