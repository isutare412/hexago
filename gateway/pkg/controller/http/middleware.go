package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/isutare412/hexago/gateway/pkg/logger"
	"go.uber.org/zap"
)

type responseWrapper struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (rw *responseWrapper) Write(body []byte) (int, error) {
	rw.body = body
	return rw.ResponseWriter.Write(body)
}

func (rw *responseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWrapper) errorMsg() string {
	if len(rw.body) == 0 || rw.statusCode < http.StatusBadRequest {
		return ""
	}

	var res errorResp
	if err := json.Unmarshal(rw.body, &res); err != nil {
		return ""
	}
	return res.ErrorMsg
}

func wrapResponseWriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		h.ServeHTTP(rw, r)
	})
}

func accessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)
		wrp := w.(*responseWrapper)

		l := logger.A().With(
			zap.Int("status", wrp.statusCode),
			zap.String("remoteAddr", r.RemoteAddr),
			zap.String("method", r.Method),
			zap.Stringer("url", r.URL),
			zap.Int64("elapsedTime", time.Since(start).Milliseconds()),
		)
		if errMsg := wrp.errorMsg(); errMsg != "" {
			l = l.With(zap.String("errorMsg", errMsg))
		}
		l.Info("Http handled")
	})
}
