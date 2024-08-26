package models

type Order struct {
	ID          string `json:"id" bson:"_id"`
	TableID     string `json:"table_id" bson:"table_id"`
	Items       []Item `json:"items" bson:"items"`
	TotalAmount int    `json:"total_amount" bson:"total_amount"`
	Status      string `json:"status" bson:"status"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
}
