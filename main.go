package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("deneme"))
	})
	http.ListenAndServe(":8080", nil)
}
