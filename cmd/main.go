package main

import (
	"encoding/json"
	"log"
	"net/http"

	"repo.analyger/internal/config"
	"repo.analyger/internal/github"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	githubClient := github.NewGithubClient()
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	result, err := githubClient.SearchGitHub(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	config.Init()

	// 현재 디렉토리에서 파일 서빙
	fs := http.FileServer(http.Dir("./wasm"))
	http.HandleFunc("/api/search", searchHandler)

	// "/static/" 경로로 접근했을 때 파일 제공
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":3334", nil))
}
