package services

import (
	"os"
	"encoding/json"

	"github.com/rhenerp/go-htmx-bootstrap/models"
)

func GetItemsInFile() []models.Item {
	items := getItemsInFile()
	return items
}

func GetItemByIdFile(id int) *models.Item {
	items := getItemsInFile()

	for i := range items {
		if items[i].Id == id {
			return &items[i] 
		}
	}

	return nil
}

func getItemsInFile() []models.Item {
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