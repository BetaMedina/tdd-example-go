package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRepository interface {
	Create(name string, email string) error
}
type memberRepository struct {
	collection *mongo.Collection
}

func (collection *memberRepository) Create(name string, email string) error {
	_, err := collection.collection.InsertOne(context.TODO(), bson.D{{Key: "name", Value: name}, {Key: "email", Value: email}})
	if err != nil {
		return err
	}
	return nil
}
func NewMemberRepository(collection *mongo.Collection) MemberRepository {
	return &memberRepository{collection}
}
