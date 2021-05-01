package main

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("index page"))

	x := r.URL.Path[1:]
	data := ""
	if len(x) > 0 {
		data = "merfaba, " + x
	} else {
		data = "index page"
	}

	w.Write([]byte(data))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("merhaba dünya"))
	//w.WriteHeader(http.StatusOK)
}

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "merhaba demir adam")
}
func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "merhaba wolverine")
}
