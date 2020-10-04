package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os/exec"
)

var (
	listenFlag      string
	pathFlag        string
	commandFlag     string
	successCodeFlag uint
	errorCodeFlag   uint
)

func init() {
	flag.StringVar(&listenFlag, "listen", ":4000", "")
	flag.StringVar(&pathFlag, "path", "/", "")
	flag.StringVar(&commandFlag, "command", "", "")
	flag.UintVar(&successCodeFlag, "success", 200, "")
	flag.UintVar(&errorCodeFlag, "error", 500, "")
	flag.Parse()
}

func main() {
	if commandFlag == "" {
		log.Fatal("missing -command")
	}

	log.Printf("Start listening on %v", listenFlag)
	http.HandleFunc(pathFlag, healthcheck(commandFlag, successCodeFlag, errorCodeFlag))
	http.ListenAndServe(listenFlag, nil)
}

func healthcheck(command string, successCode, errorCode uint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if out, err := runCommand(command); err != nil {
			log.Printf("command failed: %s", fmtOutput(out))
			w.WriteHeader(int(errorCode))
			return
		}

		w.WriteHeader(int(successCode))
	}
}

func runCommand(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	return cmd.CombinedOutput()
}

func fmtOutput(out []byte) []byte {
	if len(out) == 0 {
		return []byte("[no output]")
	}

	return bytes.ReplaceAll(out, []byte("\n"), []byte("\\n"))
}
