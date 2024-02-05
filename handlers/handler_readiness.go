package handlers

import (
	"net/http"

	"github.com/shricodev/rss-aggregator/helper"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJson(w, 200, struct{}{})
}
