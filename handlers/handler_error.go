package handlers

import (
	"net/http"

	"github.com/shricodev/rss-aggregator/helper"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithError(w, 400, "Something went wrong! Please try again later.")
}
