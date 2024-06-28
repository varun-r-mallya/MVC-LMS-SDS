package main

import (
	"net/http"
)

func main() {
	// Define handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the response directly
		http.ServeFile(w, r, "tmpl/home/index.html")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Write the response directly
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// Write the response directly
		http.ServeFile(w, r, "tmpl/login/admin.html")
	})
	// http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
	// 	// Write the response directly
	// 	http.ServeFile(w, r, "tmpl/login/client.html")
	// })

	// Start the server
	http.ListenAndServe(":8080", nil)
}
