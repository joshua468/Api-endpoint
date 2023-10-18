package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UTCTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GitHubFileURL string `json:"github_file_url"`
	GitHubRepoURL string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	slackName := params.Get("slack_name")
	track := params.Get("track")

	currentDay := time.Now().Weekday().String()
	currentTime := time.Now().UTC()
	currentTime = currentTime.Add(time.Minute * -2)

	response := Response{
		SlackName:     slackName,
		CurrentDay:    currentDay,
		UTCTime:       currentTime.Format("2006-01-02T15:04:05Z"),
		Track:         track,
		GitHubFileURL: "https://github.com/joshua468/temi-projects/edit/main/slack-project",
		GitHubRepoURL: "https://github.com/joshua468/temi-projects/slack-project",
		StatusCode:    200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", endpointHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
