package mongodb

import (
	"context"
	"go-mongodb-demo/users/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	C *mongo.Collection
}

func (m *UserModel) All() ([]models.User, error) {
	ctx := context.TODO()
	users := []models.User{}

	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, err
}
