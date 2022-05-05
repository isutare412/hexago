package entity

import "time"

type User struct {
	Id          string            `bson:"id"`
	Email       string            `bson:"email"`
	Nickname    string            `bson:"nickname"`
	GivenName   string            `bson:"given_name"`
	MiddleName  string            `bson:"middle_name"`
	FamilyName  string            `bson:"family_name"`
	Birth       time.Time         `bson:"birth"`
	DonatedFrom []*DonateRelation `bson:"donated_from"`
	DonatedTo   []*DonateRelation `bson:"donated_to"`
}

type DonateRelation struct {
	UserId    string    `bson:"user_id"`
	Nickname  string    `bson:"nickname"`
	Cents     int64     `bson:"cents"`
	Timestamp time.Time `bson:"timestamp"`
}
