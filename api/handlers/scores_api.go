package main

import (
	"clarityai/api/cache"
	"encoding/json"
	"net/http"
)

type ScoreAPI struct {
	cache *cache.Cache
}

type ScoreDetails struct {
	SecurityID string `json:"security_id"`
	Score      string `json:"score"`
}

func (api *ScoreAPI) GetScores(w http.ResponseWriter, r *http.Request) {
	securityID := r.URL.Query().Get("security_id")
	if securityID == "" {
		http.Error(w, "missing security_id parameter", http.StatusBadRequest)
		return
	}

	score, err := api.cache.Get(securityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scoreDetails := &ScoreDetails{
		SecurityID: securityID,
		Score:      score,
	}

	jsonBytes, err := json.Marshal(scoreDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
