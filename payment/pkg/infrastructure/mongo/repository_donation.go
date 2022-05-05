package mongo

import (
	"context"
	"fmt"
	"time"

	centity "github.com/isutare412/hexago/common/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) RecordDonationHistory(
	ctx context.Context,
	donatorId, donateeId string,
	cents int64,
) (err error) {
	sess, err := r.cli.StartSession()
	if err != nil {
		return fmt.Errorf("starting session: %w", err)
	}
	defer sess.EndSession(ctx)

	sCtx := mongo.NewSessionContext(ctx, sess)
	if err := sCtx.StartTransaction(); err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	var donator centity.User
	findRes := r.db.Collection(collUser).FindOne(sCtx, bson.M{"id": donatorId})
	if err := findRes.Err(); err != nil {
		return fmt.Errorf("finding donator: %w", err)
	}
	if err := findRes.Decode(&donator); err != nil {
		return fmt.Errorf("decoding donator: %w", err)
	}

	var donatee centity.User
	findRes = r.db.Collection(collUser).FindOne(sCtx, bson.M{"id": donateeId})
	if err := findRes.Err(); err != nil {
		return fmt.Errorf("finding donatee: %w", err)
	}
	if err := findRes.Decode(&donatee); err != nil {
		return fmt.Errorf("decoding donatee: %w", err)
	}

	var txTime = time.Now()

	updateRes, err := r.db.Collection(collUser).UpdateOne(sCtx,
		bson.M{"id": donatee.Id},
		bson.M{
			"$push": primitive.M{
				"donated_from": primitive.M{
					"user_id":   donator.Id,
					"nickname":  donator.Nickname,
					"cents":     cents,
					"timestamp": txTime,
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("updating donatee history: %w", err)
	}
	if updateRes.MatchedCount <= 0 {
		return fmt.Errorf("donatee not found while updating")
	}

	updateRes, err = r.db.Collection(collUser).UpdateOne(sCtx,
		bson.M{"id": donator.Id},
		bson.M{
			"$push": primitive.M{
				"donated_to": primitive.M{
					"user_id":   donatee.Id,
					"nickname":  donatee.Nickname,
					"cents":     cents,
					"timestamp": txTime,
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("updating donator history: %w", err)
	}
	if updateRes.MatchedCount <= 0 {
		return fmt.Errorf("donator not found while updating")
	}

	history := centity.DonationHistory{
		DonatorId: donator.Id,
		DonateeId: donatee.Id,
		Cents:     cents,
		Timestamp: txTime,
	}
	_, err = r.db.Collection(collDonationHistory).InsertOne(sCtx, &history)
	if err != nil {
		return fmt.Errorf("inserting donation history: %w", err)
	}

	if err := sCtx.CommitTransaction(ctx); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}
	return nil
}
