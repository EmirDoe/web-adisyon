package models

import (
	"context"
	"errors"
	"fmt"

	"webadisyon.com/db"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

type Table struct {
	ID          string `json:"id" bson:"_id"`
	TableNumber int    `json:"table_number" bson:"table_number"`
	TableStatus string `json:"table_status" bson:"table_status"`
}

func AddSingleTable(table Table) (TableID string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in AddTable")
		}
	}()

	table.ID = uuid.New().String()
	table.TableStatus = "Not Occupied"
	result, err := db.TableCollection.InsertOne(context.Background(), table)

	return fmt.Sprintf("%v", result.InsertedID), err
}

func AddTablesAtSetup(TableCount int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in AddTablesAtSetup")
		}
	}()

	for i := 1; i <= TableCount; i++ {
		table := Table{
			ID:          uuid.New().String(),
			TableNumber: i,
			TableStatus: "Not Occupied",
		}

		_, err = db.TableCollection.InsertOne(context.Background(), table)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateTable(tableID string, table Table) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in UpdateTable")
		}
	}()

	_, err = db.TableCollection.UpdateOne(context.Background(), bson.M{"_id": tableID}, bson.M{"$set": table})
	if err != nil {
		return err
	}

	return nil
}

func DeleteTable(tableID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in DeleteTable")
		}
	}()

	_, err = db.TableCollection.DeleteOne(context.Background(), bson.M{"_id": tableID})
	if err != nil {
		return err
	}

	return nil
}

func GetTables() (tables []Table, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetTables")
		}
	}()

	cursor, err := db.TableCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return tables, err
	}

	err = cursor.All(context.Background(), &tables)
	if err != nil {
		return tables, err
	}

	return tables, nil
}

func GetTableByID(tableID string) (table Table, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetTableByID")
		}
	}()

	err = db.TableCollection.FindOne(context.Background(), bson.M{"_id": tableID}).Decode(&table)
	if err != nil {
		return table, err
	}

	return table, nil
}
