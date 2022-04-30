package http

import (
	"encoding/json"
	"net/http"
)

func responseError(w http.ResponseWriter, code int, res *errorResp) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
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
