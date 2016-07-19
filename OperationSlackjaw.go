package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var otterServer string
var otterApiKey string

func main() {
	flag.StringVar(&otterServer, "otterServer", "http://ldocalhost:82", "the url to the otterserver, http://localhost:82")
	flag.StringVar(&otterApiKey, "otterApiKey", "", "the apikey that is generated in otter")
	flag.Parse()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":5080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	ok, action, appName, env := parseCommand(r.FormValue("text"))
	if !ok {
		println("not ok!")
		w.WriteHeader(http.StatusBadRequest)
	} else if action == "deploy" {
		println("ok")
		jobId := TriggerDeploy(appName, otterServer)
		response := OtterResponse{Text: "deploying " + appName + " to " + env + " with jobId: " + jobId, SlackResponseType: "in_channel"}
		json.NewEncoder(w).Encode(response)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		go FollowFeedbackOfProgressUntilCompletion(appName, jobId, env, r.FormValue("response_url"), otterServer, otterApiKey)
	}
}

type OtterResponse struct {
	Text              string `json:"text"`
	SlackResponseType string `json:"response_type"`
}
