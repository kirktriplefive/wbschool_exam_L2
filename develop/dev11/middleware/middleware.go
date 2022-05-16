package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	  start := time.Now()
	  next.ServeHTTP(w, req)
	  logrus.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
  }