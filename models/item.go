package models

type Item struct {
	ID          string `json:"id" bson:"_id"`
	MenuID      string `json:"menu_id" bson:"menu_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	ImagePath   string `json:"image" bson:"image"`
}
