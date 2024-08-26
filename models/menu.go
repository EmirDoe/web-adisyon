package models

type Menu struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Items       []Item `json:"items" bson:"items"`
}
