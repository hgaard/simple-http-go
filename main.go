package main

import (
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init()  {
	log.AddHook(logruseq.NewSeqHook("http://localhost:5341"))
}

func main() {
	log.Info("hello server")

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", ApiServer)

	log.Info("booting server")
	http.ListenAndServe(":8080", nil)

	log.Info("Closing down.. bye!")
}

