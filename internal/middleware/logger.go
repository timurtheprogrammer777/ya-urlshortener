package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	data *responseData
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.data.status = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.data.size += size
	return size, err
}

func InitLogger() error {
	logger, err := zap.NewProduction() // или zap.NewDevelopment() если хочешь более человекочитаемые логи
	if err != nil {
		return err
	}
	sugar = logger.Sugar()
	return nil
}

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.URL
		requestMethod := r.Method
		requestStart := time.Now()

		rd := &responseData{}
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			data:           rd,
		}

		h.ServeHTTP(lrw, r)

		requestDuration := time.Since(requestStart)

		sugar.Infow("request completed",
			"uri", requestURL.Path,
			"method", requestMethod,
			"status", rd.status,
			"size", rd.size,
			"duration", requestDuration,
		)
	}

	return http.HandlerFunc(logFn)
}
