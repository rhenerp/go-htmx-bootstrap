package services

import (
	"os"
	"encoding/json"

	"github.com/rhenerp/go-htmx-bootstrap/models"
)

func GetItemsInFile() []models.Item {
	var items []models.Item

	data, err := os.ReadFile("items.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		panic(err)
	}
	
	return items
}