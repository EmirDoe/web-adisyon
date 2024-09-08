package models

import (
	"context"
	"errors"

	"gopkg.in/mgo.v2/bson"
	"webadisyon.com/db"
)

type Asset struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
}

func UploadItemAsset(asset Asset, itemID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in UploadItemAsset")
		}
	}()

	_, err = db.ItemCollection.UpdateOne(context.Background(), bson.M{"_id": itemID}, bson.M{"$set": bson.M{"image": asset.Path}})
	return err

}
