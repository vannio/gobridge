package main

import (
	"fmt"
	"net/http"
)

// Hello : handles hello route
func Hello(rw http.ResponseWriter, r *http.Request) {
	// rw is the outputter, the response
	// r is the request, information that was passed into the webserver

	// Fprint converts string to a stream of bytes
	// and sends it as a response
	fmt.Fprint(rw, "Hello")
}

// Goodbye : handles goodbye route
// Using args that look more like the node code I recognise :smile:
func Goodbye(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Goodbye "+req.URL.Query().Get("name"))
}

func main() {
	// Convenience method - creates new route
	// maps handler on the DefaultServeMux
	// For any request that matches "/hello", execute HelloWorld
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/goodbye", Goodbye)

	// Routing handler: http.DefaultServeMux
	// DefaultServeMux registers routes
	// ServeMux is a type which implements the Handler
	fmt.Println("Listening for connections on port", 9000)
	http.ListenAndServe(":9000", http.DefaultServeMux)
	// New word for vocab! Mux: a term referring to a device that takes multiple inputs and forwards them into one line.
}
