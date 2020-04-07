package main

import (
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func init() {
	log.AddHook(logruseq.NewSeqHook("http://seq:5341"))
}

func main() {

	log.Info("hello server")

	n := negroni.New()
	mux := http.NewServeMux()
	n.UseHandler(mux)

	rl := NewRequestRecorder()

	n.Use(NewRequestLogger(rl))

	mux.HandleFunc("/", HelloServer)
	mux.HandleFunc("/api", ApiServer)

	ls := NewLogsServer(rl)
	mux.HandleFunc("/logs", ls.getRequestLog)

	http.ListenAndServe(":8080", n)

	log.Info("Closing down.. bye!")
}


