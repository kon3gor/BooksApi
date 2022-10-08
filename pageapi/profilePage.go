package pageapi

import (
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type ProfilePageResponse struct {
	LikedBooks []model.Book `json:"books"`
	Query string `json:"query"`
}

func ProfilePageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookIds := strings.Split(r.URL.Query().Get("books"), ",")
	query := utils.MakeLikedQuery(bookIds)
	ch := make(chan []model.Book)

	go fetchers.SearchFetcher(query, 0, ch)
	result := <-ch

	response := ProfilePageResponse {
		LikedBooks : result,
		Query : query,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

