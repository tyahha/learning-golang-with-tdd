package main

import "net/http"

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "World")
	}))
}
