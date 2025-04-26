package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.URL
		requestMethod := r.Method
		requestStart := time.Now()

		h.ServeHTTP(w, r)

		requestDuration := time.Since(requestStart)

		sugar.Infoln(
			"uri", requestURL,
			"method", requestMethod,
			"duration", requestDuration,
		)

		fmt.Println(requestURL, requestMethod, requestStart, requestDuration)
	}

	return http.HandlerFunc(logFn)
}
