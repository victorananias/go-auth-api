package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Database struct {
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	CreatedAt string
}

const ATLAS_URL = "mongodb+srv://victor:Senha123#@cluster0.lnn4l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func (database *Database) CreateUser(user User) string {
	user.Password = hashPassword(user.Password)
	user.CreatedAt = time.Now().Format("YYYY-MM-DD hh:mm:ss")
	return database.InsertInto("Users", user)
}

func (database *Database) InsertInto(collectionName string, document interface{}) (insertedId string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(ATLAS_URL))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	err = client.Connect(ctx)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	db := client.Database("Auth")

	collection := db.Collection(collectionName)

	result, err := collection.InsertOne(ctx, document)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	return result.InsertedID.(primitive.ObjectID).String()
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
		return ""
	}

	return string(hash)
}
