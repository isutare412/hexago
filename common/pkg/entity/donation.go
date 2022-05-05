package entity

import "time"

type DonationHistory struct {
	DonatorId string    `bson:"donator_id"`
	DonateeId string    `bson:"donatee_id"`
	Cents     int64     `bson:"cents"`
	Timestamp time.Time `bson:"timestamp"`
}
