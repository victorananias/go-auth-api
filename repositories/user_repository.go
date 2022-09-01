package repositories

import (
	"github.com/victorananias/go-auth-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "Users"

type UserRepository struct {
	*Repository
	collectionName string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Repository:     NewRepository(),
		collectionName: COLLECTION_NAME,
	}
}

func (repository *UserRepository) Create(user models.User) (string, error) {
	insertResult, err := repository.collection().InsertOne(repository.ctx, user)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).String(), nil
}

func (repository *UserRepository) FindUserByUsername(username string) (models.User, error) {
	where := bson.D{{Key: "username", Value: username}}
	var user models.User
	result := repository.collection().FindOne(repository.ctx, where)
	err := result.Err()
	if err != nil {
		return user, err
	}
	err = result.Decode(&user)
	return user, err
}

func (repository *UserRepository) List() ([]models.User, error) {
	users := []models.User{}
	result, err := repository.collection().Find(repository.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = result.All(repository.ctx, &users)
	return users, err
}

func (repository *UserRepository) collection() *mongo.Collection {
	return repository.db.Collection(repository.collectionName)
}
