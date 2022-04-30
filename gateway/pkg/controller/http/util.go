package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/isutare412/hexago/gateway/pkg/derror"
)

func responseError(w http.ResponseWriter, code int, res *errorResp) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

func responseDomainError(w http.ResponseWriter, err error) {
	var res errorResp

	derr := derror.Unbind(err)
	if derr == nil {
		res.Msg = err.Error()
		responseError(w, http.StatusInternalServerError, &res)
		return
	}

	for knownErr, code := range derror.DomainErrorToStatusCode {
		if errors.Is(derr, knownErr) {
			res.Msg = derr.Error()
			responseError(w, code, &res)
			return
		}
	}
	res.Msg = derr.Error()
	responseError(w, http.StatusInternalServerError, &res)
}

func responseJson(w http.ResponseWriter, res interface{}) {
	resBytes, err := json.Marshal(res)
	if err != nil {
		res := errorResp{
			Msg: "failed to marshal response",
		}
		responseError(w, http.StatusInternalServerError, &res)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(resBytes)
}
