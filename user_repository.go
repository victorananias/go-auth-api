package main

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	CreatedAt string
}

type UserRepository struct {
	*Repository
	collectionName string
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		Repository:     newRepository(),
		collectionName: "Users",
	}
}

func (repository *UserRepository) Register(user User) (err error, insertedID string) {
	user.Password = hashPassword(user.Password)
	user.CreatedAt = time.Now().Format("YYYY-MM-DD hh:mm:ss")
	findResult := repository.collection().FindOne(repository.ctx, bson.D{{"username", user.Username}})
	if err := findResult.Err(); err == nil {
		return errors.New("username already in use"), ""
	}
	insertResult, err := repository.collection().InsertOne(repository.ctx, user)
	if err != nil {
		return err, ""
	}
	return nil, insertResult.InsertedID.(primitive.ObjectID).String()
}

func (repository *UserRepository) Login(username, password string) bool {
	where := bson.D{{"username", username}}
	var user User
	result := repository.collection().FindOne(repository.ctx, where)
	if err := result.Err(); err != nil {
		return false
	}
	if err := result.Decode(&user); err != nil {
		return false
	}
	return compareHashAndPassword(user.Password, password)
}

func (repository *UserRepository) collection() *mongo.Collection {
	return repository.db.Collection(repository.collectionName)
}

func compareHashAndPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return nil == err
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}
