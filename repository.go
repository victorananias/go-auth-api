package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

const (
	defaultDb = "Auth"
	url       = "mongodb+srv://victor:Senha123#@cluster0.lnn4l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

func newRepository() *Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &Repository{
		client: client,
		db:     client.Database(defaultDb),
	}
}
