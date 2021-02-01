package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	}else if r.URL.Path == "/fizzbuzz"{
		s.counter++
		value, _ := r.URL.Query()["value"]
		w.WriteHeader(200)
		w.Write([]byte(fizzbuzz(value[0])))
	}else{
		w.WriteHeader(404)
		w.Write([]byte("404 page not found\n"))
		s.counter++
	}
}

func fizzbuzz(val string) (string)  {
	if val == "" {
		return "no value provided\n"
	}
	intVal, err := strconv.ParseInt(val,10,64)
	if err != nil{
		return "not an integer\n"
	}else{
		if intVal%3 == 0 && intVal%5 == 0 {return "fizzbuzz\n"}
		if intVal%3 == 0 {return "fizz\n"}
		if intVal%5 == 0 {return "buzz\n"}
		return fmt.Sprintf("%v\n",intVal)
	}
}
