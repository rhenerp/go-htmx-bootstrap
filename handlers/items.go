package handlers

import (
	"net/http"
	"time"

	"github.com/rhenerp/go-htmx-bootstrap/ui/modules"
	"github.com/rhenerp/go-htmx-bootstrap/services"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := services.GetItemsInFile()

	time.AfterFunc(2*time.Second, func() {
		modules.RenderSampleTable(items).Render(r.Context(), w)
	})

	time.Sleep(3 * time.Second)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {

}