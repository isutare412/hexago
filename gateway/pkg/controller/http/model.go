package http

import (
	"time"

	centity "github.com/isutare412/hexago/common/pkg/entity"
)

type errorResp struct {
	Msg string `json:"msg"`
}

type createUserReq struct {
	Id         string `json:"id" example:"isutare412"`
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
	Id         string `json:"id" example:"isutare412"`
	Email      string `json:"email" example:"foo@bar.com"`
	Nickname   string `json:"nickname" example:"redshore"`
	GivenName  string `json:"givenName"`
	MiddleName string `json:"middleName"`
	FamilyName string `json:"familyName"`
	BirthYear  int    `json:"birthYear" example:"1993"`
	BirthMonth int    `json:"birthMonth" example:"9"`
	BirthDay   int    `json:"birthDay" example:"25"`
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
}

type requestDonationReq struct {
	DonatorId string `json:"donatorId"`
	DonateeId string `json:"donateeId"`
	Cents     int64  `json:"cents" example:"150"`
}
