package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/zofy/adventure/internal/bookAPI"
)

const staticDir = "../../internal/server/static/"

var templates = map[string]*template.Template{
	"default": template.Must(template.New("").Parse("")),
	"index":   template.Must(template.ParseFiles(staticDir + "index.html")),
	"book":    template.Must(template.ParseFiles(staticDir + "book.html")),
}

type handler struct {
	story  bookAPI.Story
	tmpl   *template.Template
	pathFn func(*http.Request) string
}

// HandlerOption -
type HandlerOption func(h *handler)

// WithTemplate -
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tmpl = t
	}
}

// WithPathFn -
func WithPathFn(fn func(*http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// NewHandler -
func NewHandler(s bookAPI.Story, opts ...HandlerOption) http.Handler {
	h := handler{
		story:  s,
		tmpl:   templates["default"],
		pathFn: defaultPathFn,
	}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.story[path]; ok {
		err := h.tmpl.Execute(w, chapter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Wrong page", http.StatusNotFound)
}

// Index -
func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates["index"].Execute(w, nil)
	}
}
