package main

import ("fmt";
		"net/http")

// function to support the index handler (createdin the main function)
// writer places info/things on page, reading through the http request
func index_handler(w http.ResponseWriter, r *http.Request) {
	//FPrintf formats based on what you write
	// For example, using w (writer), write this string
	fmt.Fprintf(w, "GoLang is the bomb")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us!!")
}

func main() {
	//handlers for each endpoint
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil) //nil is none
}
