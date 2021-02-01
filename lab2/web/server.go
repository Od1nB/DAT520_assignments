package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	flag.Parse()
	server := NewServer()
	log.Fatal(http.ListenAndServe(*httpAddr, server))
}

// Server implements the web server specification found at
// lab2/README.md#web-server
type Server struct { // TODO(student): Add needed fields
	counter int
}

// NewServer returns a new Server with all required internal state initialized.
// NOTE: It should NOT start to listen on an HTTP endpoint.
func NewServer() *Server {
	s := &Server{counter: 0}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/"{
		w.WriteHeader(200)
		w.Write([]byte("Hello World!\n"))
		s.counter++
	}else if r.URL.Path == "/counter"{
		s.counter++
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("counter: %v\n",s.counter)))
	}else if r.URL.Path =="/github"{
		s.counter++
		w.WriteHeader(301)
		w.Write([]byte("<a href=\"http://www.github.com\">Moved Permanently</a>.\n\n"))
	}else{
		w.WriteHeader(404)
		w.Write([]byte("404 page not found\n"))
		s.counter++
	}

	//{"GET", "/", 200, "Hello World!\n"}
	
	//The pattern "/"" (root) should return status code 200 and the body Hello World!\n.
	// GET: "/" return 200 and "Hello World!\n"

	// /counter. A request to this pattern should return status code 200 and the current count (inclusive the current request) as the body, e.g. counter: 42\n

	//A request to the pattern /github should return status code 301 to the client with body <a href=\"http://www.github.com\">Moved Permanently</a>.\n\n

	//All other patterns should return status code 404 and the body 404 page not found\n

}
