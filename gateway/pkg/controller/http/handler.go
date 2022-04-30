package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/isutare412/hexago/gateway/pkg/core/port"
	"github.com/isutare412/hexago/gateway/pkg/logger"
)

// @Tags		User
// @Description	Create an user.
// @Router		/api/v1/users [post]
// @Param		request	body	createUserReq	true "Request to create user."
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
// @Param		email	query	string	true "Contents provider id." extensions(x-example=foo@bar.com)
// @Success		200		{object}	getUserResp
// @Failure		default	{object}	errorResp
func getUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		email := queryParams.Get("email")
		if email == "" {
			err := fmt.Errorf("need 'email' query parameter")
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		user, err := uSvc.GetByEmail(ctx, email)
		if err != nil {
			logger.S().With(
				"email", email,
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
// @Param		email	query	string	true "Contents provider id." extensions(x-example=foo@bar.com)
// @Success		200
// @Failure		default	{object}	errorResp
func deleteUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		email := queryParams.Get("email")
		if email == "" {
			err := fmt.Errorf("need 'email' query parameter")
			logger.S().Error(err)
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		if err := uSvc.DeleteByEmail(ctx, email); err != nil {
			logger.S().With(
				"email", email,
			).Error(err)
			responseDomainError(w, err)
			return
		}
	}
}
