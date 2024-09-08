package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"webadisyon.com/db"

	"github.com/golang-jwt/jwt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        string    `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	UserName  string    `bson:"username" json:"username"`
	Password  []byte    `bson:"password" json:"password" gorm:"unique"`
	Role      int       `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// User Roles
const (
	AdminRole = iota
	WaiterRole
	ChefRole
	CashierRole
)

func CreateUser(user User) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in CreateUser")
		}
	}()

	user.ID = uuid.New().String()
	result, err := db.UserCollection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	return fmt.Errorf("%v", result.InsertedID)
}

func GetUserByUserName(username string) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetUserByUserName")
		}
	}()

	err = db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return user, nil
		}
	}

	return user, err
}

func GetUserByID(id string) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetUserByID")
		}
	}()

	err = db.UserCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return user, nil
		}
	}

	return user, err
}

func GetUserByToken(token jwt.Token) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetUserByToken")
		}
	}()

	err = db.UserCollection.FindOne(context.Background(), bson.M{"token": token}).Decode(&user)

	claims := token.Claims.(*jwt.StandardClaims)
	user, err = GetUserByID(claims.Issuer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return user, nil
		}
	}

	return user, err
}
