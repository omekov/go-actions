package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Starting server for github"))
	})
	http.ListenAndServe(":80", nil)
}
