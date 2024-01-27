package main

import (
	"io"
	"net/http"
)

func main() {
	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// consider cors from localhost:5173
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// write hi to body
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, "rs. 20")
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", rootHandler)
	if err != nil {
		return
	}
}
