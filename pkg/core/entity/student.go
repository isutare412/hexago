package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id         primitive.ObjectID `bson:"_id"`
	GivenName  string             `bson:"given_name"`
	MiddleName string             `bson:"middle_name"`
	FamilyName string             `bson:"family_name"`
	Birth      time.Time          `bson:"birth"`
}
