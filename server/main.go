package main

import (
	"fmt"
	"net/http"
	"simple-http-server/logger"

	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func init() {
	seqURL := logger.GetEnv("SEQ_URL", "http://localhost:5341")
	fmt.Printf("Logging to SEQ_URL '%s'\n", seqURL)
	log.AddHook(logruseq.NewSeqHook(seqURL))
}

func main() {

	log.Info("hello server")

	n := negroni.New()
	mux := http.NewServeMux()
	n.UseHandler(mux)

	rl := logger.NewRequestRecorder()

	n.Use(logger.NewRequestLogger(rl))

	mux.HandleFunc("/", HelloServer)
	mux.HandleFunc("/api", ApiServer)

	ls := NewLogsServer(rl)
	mux.HandleFunc("/logs", ls.getRequestLog)

	http.ListenAndServe(":8080", n)

	log.Info("Closing down.. bye!")
}
