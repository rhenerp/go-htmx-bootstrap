package handlers

import (
	"net/http"

	"github.com/rhenerp/go-htmx-bootstrap/ui/layouts"
)

func Home(w http.ResponseWriter, r *http.Request) {
	layouts.BaseLayout().Render(r.Context(), w)
}