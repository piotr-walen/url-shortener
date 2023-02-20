package utils

import (
	"log"
	"net/http"
)

type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func AttachLogger(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.String()
		method := r.Method
		record := &LogRecord{
			ResponseWriter: w,
		}

		f.ServeHTTP(record, r)

		log.Println(method, record.status, uri)
	}
}
