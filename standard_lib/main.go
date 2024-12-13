package main

import (
	"net/http"
)

func main() {
	// Create a multiplexer to interface w/ http requests and match them with
	// URI patterns
	mux := http.NewServeMux()

	// Register URI paths and define which HTTP Handler to execute for the URI
	mux.Handle("/", &homeHandler{})

	// Call the ListenAndServe function to start your sever, assigning the mux
	// created above to handle requests
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// Create URI path structs and implement ServeHTTP
type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page\n"))
}
