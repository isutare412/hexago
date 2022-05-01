package centity

import "time"

type User struct {
	Id         string    `bson:"id"`
	Email      string    `bson:"email"`
	Nickname   string    `bson:"nickname"`
	GivenName  string    `bson:"given_name"`
	MiddleName string    `bson:"middle_name"`
	FamilyName string    `bson:"family_name"`
	Birth      time.Time `bson:"birth"`
}
