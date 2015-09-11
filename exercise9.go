// https://tour.golang.org/methods/14

// Exercise: HTTP Handlers
// Implement the following types and define ServeHTTP methods on them. Register them to handle specific paths in your web server.
//
// type String string
//
// type Struct struct {
//     Greeting string
//     Punct    string
//     Who      string
// }
// For example, you should be able to register handlers using:
//
// http.Handle("/string", String("I'm a frayed knot."))
// http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
// Note: This example won't run through the web-based tour user interface. To try writing web servers you may want to Install Go.

package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func Exercise9() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot!"))
	http.Handle("/struct", &Struct{"Hola", ":", "Gopher"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
