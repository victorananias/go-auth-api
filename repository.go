package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client   *mongo.Client
	db       *mongo.Database
	ctx      context.Context
	settings Settings
}

const defaultDb = "Auth"

func newRepository() *Repository {
	repository := &Repository{}
	settings, err := newSettings()

	if err != nil {
		log.Fatalf(err.Error())
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(settings.Db.ConnectionString))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	repository.client = client
	repository.db = client.Database(defaultDb)
	repository.settings = *settings
	return repository
}
