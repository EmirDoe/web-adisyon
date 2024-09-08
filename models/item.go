package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"webadisyon.com/db"
)

type Item struct {
	ID          string `json:"id" bson:"_id"`
	MenuID      int    `json:"menu_id" bson:"menu_id"`
	Category    int    `json:"category" bson:"category"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	ImagePath   string `json:"image" bson:"image"`
}

type ItemOnOrder struct {
	ItemID   string `json:"item_id" bson:"item_id"`
	Price    int    `json:"price" bson:"price"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

type ItemCategory struct {
	ID   int    `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func AddItemCategory(category ItemCategory) (CategoryID string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in AddItemCategory")
		}
	}()

	//check if collection is empty
	count, err := db.ItemCategoryCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return "", err
	}
	if count == 0 {
		category.ID = 1
		result, err := db.ItemCategoryCollection.InsertOne(context.Background(), category)
		return fmt.Sprintf("%v", result.InsertedID), err
	}

	// find last items id and increment it
	var lastItem ItemCategory
	err = db.ItemCategoryCollection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&lastItem)
	if err != nil {
		return "", err
	}
	category.ID = lastItem.ID + 1

	if category.Name == "" {
		return "", errors.New("Missing fields")
	}

	result, err := db.ItemCategoryCollection.InsertOne(context.Background(), category)

	return fmt.Sprintf("%v", result.InsertedID), err
}

func UpdateItemCategory(categoryID int, category ItemCategory) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in UpdateItemCategory")
		}
	}()

	_, err = db.ItemCategoryCollection.UpdateOne(context.Background(), map[string]interface{}{"_id": categoryID}, map[string]interface{}{"$set": category})
	if err != nil {
		return err
	}

	return nil
}

func DeleteItemCategory(categoryID int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in DeleteItemCategory")
		}
	}()

	_, err = db.ItemCategoryCollection.DeleteOne(context.Background(), map[string]interface{}{"_id": categoryID})
	if err != nil {
		return err
	}

	return nil
}

func GetItemCategory(categoryID int) (category ItemCategory, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetItemCategory")
		}
	}()

	err = db.ItemCategoryCollection.FindOne(context.Background(), map[string]interface{}{"_id": categoryID}).Decode(&category)
	if err != nil {
		return ItemCategory{}, err
	}

	return category, nil
}

func GetItemCategories() (categories []ItemCategory, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error in GetItemCategories")
		}
	}()

	cursor, err := db.ItemCategoryCollection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func AddItem(item Item) (ItemID string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in AddItem")
		}
	}()

	item.ID = uuid.New().String()
	if item.MenuID == 0 || item.Category == 0 || item.Name == "" || item.Description == "" || item.Price == 0 || item.ImagePath == "" {
		return "", errors.New("Missing fields")
	}

	result, err := db.ItemCollection.InsertOne(context.Background(), item)

	return fmt.Sprintf("%v", result.InsertedID), err
}

func UpdateItem(itemID string, item Item) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in UpdateItem")
		}
	}()

	_, err = db.ItemCollection.UpdateOne(context.Background(), map[string]interface{}{"_id": itemID}, map[string]interface{}{"$set": item})
	if err != nil {
		return err
	}

	return nil
}

func DeleteItem(itemID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in DeleteItem")
		}
	}()

	_, err = db.ItemCollection.DeleteOne(context.Background(), map[string]interface{}{"_id": itemID})
	if err != nil {
		return err
	}

	return nil
}

func GetItems() (items []Item, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetItems")
		}
	}()

	cursor, err := db.ItemCollection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemByCategory(category int) (items []Item, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetItemByCategory")
		}
	}()

	cursor, err := db.ItemCollection.Find(context.Background(), map[string]interface{}{"category": category})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemByID(itemID string) (item Item, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetItemByID")
		}
	}()

	err = db.ItemCollection.FindOne(context.Background(), map[string]interface{}{"_id": itemID}).Decode(&item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}
