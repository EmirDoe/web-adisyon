package models

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"webadisyon.com/db"
)

type Order struct {
	ID        string        `json:"id" bson:"_id"`
	TableID   string        `json:"table_id" bson:"table_id"`
	Items     []ItemOnOrder `json:"items" bson:"items"`
	Total     int           `json:"total_amount" bson:"total_amount"`
	Status    string        `json:"status" bson:"status"`
	CreatedBy string        `json:"created_by" bson:"created_by"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	Actions   []Action      `json:"actions" bson:"actions"`
}

type Action struct {
	UserID     string        `json:"user_id" bson:"user_id"`
	Items      []ItemOnOrder `json:"items" bson:"items"`
	ActionType string        `json:"action_type" bson:"action_type"`
	Timestamp  time.Time     `json:"timestamp" bson:"timestamp"`
}

func CreateOrder(order Order) (OrderID string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	order.ID = uuid.New().String()
	order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	result, err := db.OrderCollection.InsertOne(context.Background(), order)

	return fmt.Sprintf("%v", result.InsertedID), err
}

func CalculateTotal(order Order) (total int) {
	for _, item := range order.Items {
		total += item.Price * item.Quantity
	}
	return total
}

func AddAction(orderID string, action Action) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID}, bson.M{"$push": bson.M{"actions": action}})

	//get every items quantity from order and increase them by action quantity if they exist. if they dont exist add them to order check if they exist before adding,
	// first search collection for item if it exists increase quantity if it doesnt exist add it to order
	for _, item := range action.Items {
		var itemOnOrder ItemOnOrder
		err = db.OrderCollection.FindOne(context.Background(), bson.M{"_id": orderID, "items.item_id": item.ItemID}).Decode(&itemOnOrder)
		if err != nil {
			_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID}, bson.M{"$push": bson.M{"items": item}})
		} else {
			_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID, "items.item_id": item.ItemID}, bson.M{"$inc": bson.M{"items.$.quantity": item.Quantity}})
		}
	}

	return err
}

func RemoveAction(orderID string, action Action) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID}, bson.M{"$push": bson.M{"actions": action}})

	//get every items quantity from order and decrease them by action quantity if its zero remove it from order
	for _, item := range action.Items {
		_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID, "items.item_id": item.ItemID}, bson.M{"$inc": bson.M{"items.$.quantity": -item.Quantity}})
		_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID, "items.item_id": item.ItemID, "items.quantity": 0}, bson.M{"$pull": bson.M{"items": bson.M{"item_id": item.ItemID}}})
	}

	//recalculate total
	_, err = db.OrderCollection.UpdateOne(context.Background(), bson.M{"_id": orderID}, bson.M{"$set": bson.M{"total_amount": CalculateTotal(Order{ID: orderID})}})

	return err

}
