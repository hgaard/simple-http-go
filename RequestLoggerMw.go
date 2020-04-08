package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestLogger struct {
	recorder *RequestRecorder
}

func NewRequestLogger(r *RequestRecorder) *RequestLogger {
	rl := &RequestLogger{recorder: r}
	return rl
}

func (rl *RequestLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.WithFields(log.Fields{"Path": r.URL.Path, "Query": r.URL.RawQuery}).Info("Executing request pipeline")

	t := time.Now()
	next(rw, r)

	// look at timer
	elapsed := time.Since(t)
	rl.recorder.add(RequestRecord{r.RequestURI, elapsed})
	log.WithFields(log.Fields{"Elapsed time": elapsed}).Info("Executed request in %v", elapsed)
}
