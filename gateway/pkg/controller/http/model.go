package http

import (
	"time"

	centity "github.com/isutare412/hexago/common/pkg/entity"
)

type errorResp struct {
	ErrorMsg string `json:"errorMsg"`
}

type createUserReq struct {
	Id         string `json:"id" example:"id001"`
	Email      string `json:"email" example:"foo@bar.com"`
	Nickname   string `json:"nickname" example:"redshore"`
	GivenName  string `json:"givenName"`
	MiddleName string `json:"middleName"`
	FamilyName string `json:"familyName"`
	BirthYear  int    `json:"birthYear" example:"1993"`
	BirthMonth int    `json:"birthMonth" example:"9"`
	BirthDay   int    `json:"birthDay" example:"25"`
}

func (r *createUserReq) IntoUser() *centity.User {
	birth := time.Date(
		r.BirthYear, time.Month(r.BirthMonth), r.BirthDay, 0, 0, 0, 0,
		time.UTC)

	return &centity.User{
		Id:         r.Id,
		Email:      r.Email,
		Nickname:   r.Nickname,
		GivenName:  r.GivenName,
		MiddleName: r.MiddleName,
		FamilyName: r.FamilyName,
		Birth:      birth,
	}
}

type getUserResp struct {
	Id          string            `json:"id" example:"id001"`
	Email       string            `json:"email" example:"foo@bar.com"`
	Nickname    string            `json:"nickname" example:"redshore"`
	GivenName   string            `json:"givenName"`
	MiddleName  string            `json:"middleName"`
	FamilyName  string            `json:"familyName"`
	BirthYear   int               `json:"birthYear" example:"1993"`
	BirthMonth  int               `json:"birthMonth" example:"9"`
	BirthDay    int               `json:"birthDay" example:"25"`
	DonatedFrom []*donateRelation `json:"donatedFrom"`
	DonatedTo   []*donateRelation `json:"donatedTo"`
}

type donateRelation struct {
	UserId    string    `json:"userId" example:"id001"`
	Nickname  string    `json:"nickname" example:"redshore"`
	Cents     int64     `json:"cents" example:"120"`
	Timestamp time.Time `json:"timestamp" example:"2022-05-05T06:22:40.328Z"`
}

func (dr *donateRelation) FromDonateRelation(src *centity.DonateRelation) {
	dr.UserId = src.UserId
	dr.Nickname = src.Nickname
	dr.Cents = src.Cents
	dr.Timestamp = src.Timestamp
}

func (r *getUserResp) FromUser(user *centity.User) {
	r.Id = user.Id
	r.Email = user.Email
	r.Nickname = user.Nickname
	r.GivenName = user.GivenName
	r.MiddleName = user.MiddleName
	r.FamilyName = user.FamilyName
	r.BirthYear = user.Birth.Year()
	r.BirthMonth = int(user.Birth.Month())
	r.BirthDay = user.Birth.Day()

	r.DonatedFrom = make([]*donateRelation, 0, len(user.DonatedFrom))
	for _, drEntity := range user.DonatedFrom {
		var dr donateRelation
		dr.FromDonateRelation(drEntity)
		r.DonatedFrom = append(r.DonatedFrom, &dr)
	}

	r.DonatedTo = make([]*donateRelation, 0, len(user.DonatedTo))
	for _, drEntity := range user.DonatedTo {
		var dr donateRelation
		dr.FromDonateRelation(drEntity)
		r.DonatedTo = append(r.DonatedTo, &dr)
	}
}

type requestDonationReq struct {
	DonatorId string `json:"donatorId" example:"id001"`
	DonateeId string `json:"donateeId" example:"id002"`
	Cents     int64  `json:"cents" example:"150"`
}
