package pageapi

import (
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SearchResultPageResponse struct {
	SearchResult []model.Book `json:"search_result"`
	Query string `json:"query"`
}

func SearchResultPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query := r.URL.Query().Get("q")
	ch := make(chan []model.Book)

	go fetchers.SearchFetcher(query, 15, ch)
	result := <-ch

	response := SearchResultPageResponse {
		SearchResult : result,
		Query : query,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

