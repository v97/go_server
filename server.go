package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving on port 8000")
}

func page_1_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page 1")
}

func main() {
	fmt.Println("Starting server")

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/page_1/", page_1_handler)

	http.ListenAndServe(":8000", nil)
}