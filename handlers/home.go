package handlers

import (
	"net/http"
	"github.com/rhenerp/go-htmx-bootstrap/ui/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.RenderHomePage().Render(r.Context(), w)
}