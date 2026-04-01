package console

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/libdefinite/definite/internal/ctl/console/templates"
)

//go:embed static/output.css static/output.js static/favicon.svg
var staticFiles embed.FS

// Service serves the web UI for the definite control plane.
type Service struct {
	port int
}

// New creates a web Console listening on the given port.
func New(port int) *Service {
	return &Service{port: port}
}

// Start begins serving the web console and blocks until the server stops.
func (w *Service) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", w.port), w.routes())
}

func (w *Service) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", handleStatic())
	mux.HandleFunc("/health", w.handleHealth)
	mux.HandleFunc("/{$}", w.handleHome)
	mux.HandleFunc("/", w.handleNotFound)
	return mux
}

func (w *Service) handleHome(rw http.ResponseWriter, r *http.Request) {
	if err := templates.Home().Render(r.Context(), rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (w *Service) handleNotFound(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	if err := templates.NotFound().Render(r.Context(), rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (w *Service) handleHealth(rw http.ResponseWriter, r *http.Request) {
	if err := templates.HealthStatus().Render(r.Context(), rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handleStatic() http.Handler {
	sub, _ := fs.Sub(staticFiles, "static")
	return http.StripPrefix("/static/", http.FileServer(http.FS(sub)))
}
