package repo

import (
	"context"
	"fmt"

	"github.com/isutare412/hexago/gateway/pkg/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mdb *MongoDB) InsertUser(
	ctx context.Context,
	user *entity.User,
) error {
	coll := mdb.db.Collection(collUser)

	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (mdb *MongoDB) FindUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	coll := mdb.db.Collection(collUser)

	filter := bson.D{primitive.E{Key: "email", Value: email}}
	res := coll.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("finding user: %w", err)
	}

	var user entity.User
	if err := res.Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding user: %w", err)
	}
	return &user, nil
}

func (mdb *MongoDB) DeleteUserByEmail(
	ctx context.Context,
	email string,
) error {
	coll := mdb.db.Collection(collUser)

	filter := bson.D{primitive.E{Key: "email", Value: email}}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	if res.DeletedCount <= 0 {
		return fmt.Errorf("no such user with email[%s]", email)
	}

	return nil
}
