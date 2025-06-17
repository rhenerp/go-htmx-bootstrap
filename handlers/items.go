package handlers

import (
	"net/http"

	"github.com/rhenerp/go-htmx-bootstrap/models"
	"github.com/rhenerp/go-htmx-bootstrap/ui/modules"
	"time"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := []models.Item{
		{
			Name: "Rhen",
			Age: 28,
			Options: false,
		},
		{
			Name: "Hikaru Nagi",
			Age: 29,
			Options: true,
		},
	}
	
	time.AfterFunc(2*time.Second, func() {
		modules.RenderSampleTable(items).Render(r.Context(), w)
	})

	time.Sleep(3 * time.Second)
}