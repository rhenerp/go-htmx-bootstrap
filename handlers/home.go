package handlers

import (
	"net/http"

	"github.com/rhenerp/go-htmx-bootstrap/ui/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.RenderHomePage("Its ok!").Render(r.Context(), w)
}