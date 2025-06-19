package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rhenerp/go-htmx-bootstrap/services"
	"github.com/rhenerp/go-htmx-bootstrap/ui/modules"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := services.GetItemsInFile()

	time.AfterFunc(2*time.Second, func() {
		modules.RenderSampleTable(items).Render(r.Context(), w)
	})

	time.Sleep(3 * time.Second)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	itemIdParam := chi.URLParam(r, "itemId")
	itemId, err := strconv.Atoi(itemIdParam)
	if err != nil {
		http.Error(w, "Invalid item Id", http.StatusBadRequest)
		return
	}

	item := services.GetItemByIdFile(itemId)
	if item == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	} 
	modules.RenderSampleTabs(item).Render(r.Context(), w)
}