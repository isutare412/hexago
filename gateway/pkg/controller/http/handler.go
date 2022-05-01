package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/isutare412/hexago/gateway/pkg/logger"
	"github.com/isutare412/hexago/gateway/pkg/port"
)

// @Tags		User
// @Description	Create an user.
// @Router		/api/v1/users [post]
// @Param		request	body	createUserReq	true	"Request to create user."
// @Success		200
// @Failure		default	{object}	errorResp
func createUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqBytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.S().Error(err)
			responseError(w, http.StatusInternalServerError, &errorResp{
				Msg: "failed to read request body",
			})
			return
		}

		var req createUserReq
		if err := json.Unmarshal(reqBytes, &req); err != nil {
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: "request body does not follow required format",
			})
			return
		}

		user := req.IntoUser()
		if err := uSvc.Create(ctx, user); err != nil {
			logger.S().With(
				"id", user.Id,
				"email", user.Email,
			).Error(err)
			responseDomainError(w, err)
			return
		}
	}
}

// @Tags		User
// @Description	Get an user.
// @Router		/api/v1/users [get]
// @Param		id	query	string	true	"Id of user."	extensions(x-example=isutare412)
// @Success		200		{object}	getUserResp
// @Failure		default	{object}	errorResp
func getUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		id := queryParams.Get("id")
		if id == "" {
			err := fmt.Errorf("need 'id' query parameter")
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		user, err := uSvc.GetById(ctx, id)
		if err != nil {
			logger.S().With(
				"id", id,
			).Error(err)
			responseDomainError(w, err)
			return
		}

		var res getUserResp
		res.FromUser(user)
		responseJson(w, &res)
	}
}

// @Tags		User
// @Description	Delete an user.
// @Router		/api/v1/users [delete]
// @Param		id	query	string	true	"Id of user."	extensions(x-example=isutare412)
// @Success		200
// @Failure		default	{object}	errorResp
func deleteUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		id := queryParams.Get("id")
		if id == "" {
			err := fmt.Errorf("need 'id' query parameter")
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		if err := uSvc.DeleteById(ctx, id); err != nil {
			logger.S().With(
				"id", id,
			).Error(err)
			responseDomainError(w, err)
			return
		}
	}
}

// @Tags		Donation
// @Description	Request donation.
// @Router		/api/v1/donations [post]
// @Param		request	body	requestDonationReq	true	"Donation request."
// @Success		200
// @Failure		default	{object}	errorResp
func requestDonation(dSvc port.DonationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqBytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.S().Error(err)
			responseError(w, http.StatusInternalServerError, &errorResp{
				Msg: "failed to read request body",
			})
			return
		}

		var req requestDonationReq
		if err := json.Unmarshal(reqBytes, &req); err != nil {
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: "request body does not follow required format",
			})
			return
		}

		err = dSvc.RequestDonation(
			ctx, req.DonatorId, req.DonateeId, req.Cents)
		if err != nil {
			logger.S().With(
				"donatorId", req.DonatorId,
				"donateeId", req.DonateeId,
				"cents", req.Cents,
			).Error(err)
			responseDomainError(w, err)
			return
		}
	}
}
