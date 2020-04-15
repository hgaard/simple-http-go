package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	logrus.WithFields(logrus.Fields{"requestId": r.URL.Path}).Info("Serving request")
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}
