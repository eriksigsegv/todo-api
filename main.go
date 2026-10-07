package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", mux)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
