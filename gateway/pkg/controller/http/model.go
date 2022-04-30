package http

import (
	"time"

	"github.com/isutare412/hexago/gateway/pkg/core/entity"
)

type errorResp struct {
	Msg string `json:"msg"`
}

type createUserReq struct {
	Email      string `json:"email" example:"foo@bar.com"`
	Nickname   string `json:"nickname"`
	GivenName  string `json:"givenName"`
	MiddleName string `json:"middleName"`
	FamilyName string `json:"familyName"`
	BirthYear  int    `json:"birthYear" example:"1993"`
	BirthMonth int    `json:"birthMonth" example:"9"`
	BirthDay   int    `json:"birthDay" example:"25"`
}

func (r *createUserReq) IntoUser() *entity.User {
	birth := time.Date(
		r.BirthYear, time.Month(r.BirthMonth), r.BirthDay, 0, 0, 0, 0,
		time.UTC)

	return &entity.User{
		Email:      r.Email,
		Nickname:   r.Nickname,
		GivenName:  r.GivenName,
		MiddleName: r.MiddleName,
		FamilyName: r.FamilyName,
		Birth:      birth,
	}
}

type getUserResp struct {
	Email      string `json:"email" example:"foo@bar.com"`
	Nickname   string `json:"nickname"`
	GivenName  string `json:"givenName"`
	MiddleName string `json:"middleName"`
	FamilyName string `json:"familyName"`
	BirthYear  int    `json:"birthYear" example:"1993"`
	BirthMonth int    `json:"birthMonth" example:"9"`
	BirthDay   int    `json:"birthDay" example:"25"`
}

func (r *getUserResp) FromUser(user *entity.User) {
	r.Email = user.Email
	r.Nickname = user.Nickname
	r.GivenName = user.GivenName
	r.MiddleName = user.MiddleName
	r.FamilyName = user.FamilyName
	r.BirthYear = user.Birth.Year()
	r.BirthMonth = int(user.Birth.Month())
	r.BirthDay = user.Birth.Day()
}
