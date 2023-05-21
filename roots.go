package celeritas

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (c *Celeritas) routes() http.Handler {
	mux := chi.NewRouter()
}
