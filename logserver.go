package main

import "net/http"

type LogsServer struct {
	recorder *RequestRecorder
}

func NewLogsServer(r *RequestRecorder) *LogsServer{
	s := &LogsServer{ recorder:r}
	return s
}

func (l *LogsServer) getRequestLog(w http.ResponseWriter, r *http.Request) {
	js, err := l.recorder.toJson()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


