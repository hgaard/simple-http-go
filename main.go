package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// https://stackoverflow.com/a/40326580
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	seqURL := getEnv("SEQ_URL", "http://localhost:5341")
	fmt.Printf("Logging to SEQ_URL '%s'\n", seqURL)
	log.AddHook(logruseq.NewSeqHook(seqURL))
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
