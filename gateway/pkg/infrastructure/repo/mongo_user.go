package repo

import (
	"context"
	"errors"
	"fmt"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"github.com/isutare412/hexago/gateway/pkg/derror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mdb *MongoDB) InsertUser(
	ctx context.Context,
	user *centity.User,
) error {
	coll := mdb.db.Collection(collUser)

	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (mdb *MongoDB) FindUserById(
	ctx context.Context,
	id string,
) (*centity.User, error) {
	coll := mdb.db.Collection(collUser)

	filter := bson.D{primitive.E{Key: "id", Value: id}}
	res := coll.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		err = fmt.Errorf("finding user: %w", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, derror.Bind(
				err,
				fmt.Errorf("%w: no such user with id[%s]",
					derror.ErrUserNotFound, id),
			)
		}
		return nil, err
	}

	var user centity.User
	if err := res.Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding user: %w", err)
	}
	return &user, nil
}

func (mdb *MongoDB) DeleteUserById(
	ctx context.Context,
	id string,
) error {
	coll := mdb.db.Collection(collUser)

	filter := bson.D{primitive.E{Key: "id", Value: id}}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("deleting user: %w", err)
	}
	if res.DeletedCount <= 0 {
		err := fmt.Errorf("no such user with id[%s]", id)
		return derror.Bind(
			err,
			fmt.Errorf("%w: %v", derror.ErrUserNotFound, err),
		)
	}
	return nil
}
