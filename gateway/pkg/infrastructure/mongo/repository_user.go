package mongo

import (
	"context"
	"fmt"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"github.com/isutare412/hexago/gateway/pkg/derror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) InsertUser(
	ctx context.Context,
	user *centity.User,
) error {
	coll := r.db.Collection(collUser)

	_, err := coll.InsertOne(ctx, user)
	if isErrDupKey(err) {
		return derror.Bind(
			err,
			fmt.Errorf("%w: %v", derror.ErrDuplicateKey, err),
		)
	} else if err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (r *Repository) FindUserById(
	ctx context.Context,
	id string,
) (*centity.User, error) {
	coll := r.db.Collection(collUser)

	filter := bson.D{primitive.E{Key: "id", Value: id}}
	res := coll.FindOne(ctx, filter)
	if err := res.Err(); isErrNotFound(err) {
		return nil, derror.Bind(
			err,
			fmt.Errorf("%w: no such user with id[%s]",
				derror.ErrUserNotFound, id),
		)
	} else if err != nil {
		return nil, err
	}

	var user centity.User
	if err := res.Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding user: %w", err)
	}
	return &user, nil
}

func (r *Repository) DeleteUserById(
	ctx context.Context,
	id string,
) error {
	coll := r.db.Collection(collUser)

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
