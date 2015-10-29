package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server started!")

	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// log.Fatal(http.ListenAndServe(":9999", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Helll??s, %q", html.EscapeString(r.URL.Path))
	// })))
	log.Fatal(http.ListenAndServe(":9999", http.FileServer(assetFS())))

	fmt.Println("server end!")
}
