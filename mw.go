package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func MyMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	log.WithFields(log.Fields{"Path": req.URL.Path, "Query": req.URL.RawQuery}).Info("Executing request pipeline")
	t := time.Now()
	next(rw, req)
	// do some stuff after
	elapsed := time.Since(t)
	log.WithFields(log.Fields{"Elapsed time": elapsed}).Info("Executed request in %v", elapsed)
}
