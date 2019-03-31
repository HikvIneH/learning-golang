package home

import "net/http"

// Index to handle home router
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
