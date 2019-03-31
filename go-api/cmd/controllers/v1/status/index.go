package status

import (
	"net/http"
)

// Index for status
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V1 status is live!"))
}
