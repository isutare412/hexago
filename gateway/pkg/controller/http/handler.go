package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/isutare412/hexago/gateway/pkg/core/port"
	"github.com/isutare412/hexago/gateway/pkg/logger"
)

func createUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqBytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.L().Error(err.Error())
			responseError(w, http.StatusInternalServerError, &errorResp{
				Msg: "failed to read request body",
			})
			return
		}

		var req createUserReq
		if err := json.Unmarshal(reqBytes, &req); err != nil {
			logger.L().Error(err.Error())
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: "request body does not follow required format",
			})
			return
		}

		user := req.IntoUser()
		if err := uSvc.Create(ctx, user); err != nil {
			logger.L().Error(err.Error())
			responseError(w, http.StatusInternalServerError, &errorResp{
				Msg: err.Error(),
			})
			return
		}
	}
}

func getUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		email := queryParams.Get("email")
		if email == "" {
			err := fmt.Errorf("need 'email' query parameter")
			logger.L().Error(err.Error())
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		user, err := uSvc.GetByEmail(ctx, email)
		if err != nil {
			logger.L().Error(err.Error())
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		var res getUserResp
		res.FromUser(user)
		responseJson(w, &res)
	}
}

func deleteUser(uSvc port.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		email := queryParams.Get("email")
		if email == "" {
			err := fmt.Errorf("need 'email' query parameter")
			logger.L().Error(err.Error())
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}

		if err := uSvc.DeleteByEmail(ctx, email); err != nil {
			logger.L().Error(err.Error())
			responseError(w, http.StatusBadRequest, &errorResp{
				Msg: err.Error(),
			})
			return
		}
	}
}
