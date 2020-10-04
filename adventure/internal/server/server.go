package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/zofy/adventure/internal/bookAPI"
)

// Server -
type Server struct {
	port      int
	router    *http.ServeMux
	templates map[string]*template.Template
}

// New -
func New(port int) *Server {
	return &Server{
		port:   port,
		router: http.NewServeMux(),
	}
}

func (s *Server) routes(story bookAPI.Story) {
	bookH := NewHandler(story, WithTemplate(templates["book"]), WithPathFn(defaultPathFn))
	s.router.Handle("/", Index())
	s.router.Handle("/story/", bookH)
}

// Start -
func (s *Server) Start(story bookAPI.Story) {
	fmt.Printf("Starting server at port: %d\n", s.port)
	s.routes(story)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
