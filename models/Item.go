package models

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Options bool `json:"options"`
	Additional map[string]string `json:"additional"`
}
