package main

import (
	"net/http"
	"simple-http-server/logger"
)

type LogsServer struct {
	recorder *logger.RequestRecorder
}

func NewLogsServer(r *logger.RequestRecorder) *LogsServer{
	s := &LogsServer{ recorder:r}
	return s
}

func (l *LogsServer) getRequestLog(w http.ResponseWriter, r *http.Request) {
	js, err := l.recorder.ToJson()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


