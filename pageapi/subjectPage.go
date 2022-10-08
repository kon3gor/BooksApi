package pageapi

import (
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SubjectPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	subjectId := ps.ByName("subjectId")
	ch := make(chan model.Subject)

	go fetchers.SubjectFetcher(subjectId, 15, ch)
	result := <-ch

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
