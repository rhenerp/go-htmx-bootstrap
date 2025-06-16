package handlers

import (
	"net/http"

	"github.com/rhenerp/go-htmx-bootstrap/ui/pages"
	"github.com/rhenerp/go-htmx-bootstrap/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	items := []models.Item{
		{
			Name: "Rhen",
			Age: 28,
			Options: false,
		},
		{
			Name: "sdddsp",
			Age: 29,
			Options: true,
		},
	}
	pages.RenderHomePage(items).Render(r.Context(), w)
}