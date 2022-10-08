package pageapi

import (
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/utils"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var subjects = []string {
    "award:nebula_award=novel",
    "award:hugo_award=novel",
    "collectionid:cyoa1",
    "collectionid:cyoa4",
    "collectionid:goosebumps6",
    "award:euler_book_prize",
}

type ListsPageResponse struct {
	FirstCarousel []model.Book `json:"first"`
	SecondCarousel []model.Book `json:"second"`
	ThirdCarousel []model.Book `json:"third"`
}

func ListsPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inds := utils.GenerateThreeRandoms(0, len(subjects))
	ch := make(chan model.Subject)
	for _, ind := range inds {
		go fetchers.SubjectFetcher(subjects[ind], 5, ch)
	}

	result := make([]model.Subject, 3)
	for i := range inds {
		result[i] = <-ch 
	}

	response := ListsPageResponse {
		FirstCarousel : result[0].Books,
		SecondCarousel : result[1].Books,
		ThirdCarousel : result[2].Books,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

