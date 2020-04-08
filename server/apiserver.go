package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Record struct {
	Key   string
	Value string
}

func ApiServer(w http.ResponseWriter, r *http.Request) {

	logrus.WithFields(logrus.Fields{"Path": r.URL.Path, "Query": r.URL.RawQuery}).Info("Serving API request")

	key, ok := r.URL.Query()["key"]

	if !ok || len(key[0]) < 1 {
		logrus.Warn("Key param was not found")
	}

	value, ok := r.URL.Query()["value"]

	if !ok || len(value[0]) < 1 {
		logrus.Warn("value param was not found")
	}

	record := Record{key[0], value[0]}

	js, err := json.Marshal(record)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

