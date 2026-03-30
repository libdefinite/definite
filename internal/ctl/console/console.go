package console

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/libdefinite/definite/internal/ctl/console/templates"
)

//go:embed static/output.css
var staticFiles embed.FS

// Console serves the web UI for the definite control plane.
type Console struct {
	port int
}

// NewConsole creates a web Console listening on the given port.
func NewConsole(port int) *Console {
	return &Console{port: port}
}

// Start begins serving the web console and blocks until the server stops.
func (w *Console) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", w.port), w.routes())
}

func (w *Console) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", handleStatic())
	mux.HandleFunc("/", w.handleHome)
	return mux
}

func (w *Console) handleHome(rw http.ResponseWriter, r *http.Request) {
	if err := templates.Home().Render(r.Context(), rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func handleStatic() http.Handler {
	sub, _ := fs.Sub(staticFiles, "static")
	return http.StripPrefix("/static/", http.FileServer(http.FS(sub)))
}
