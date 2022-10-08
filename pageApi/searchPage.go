package pageapi

import(
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
)

type pageResponse struct {
	TrendingCarousel []model.Book `json:"carousel"`
	Subjects []model.Subject `json:"subjects"`
}

func SearchPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	trendingNow, err := fetchers.TrendingFetcher()
	if err != nil {
		log.Fatalln(err)
	}

	response := pageResponse {
		TrendingCarousel : trendingNow,
		Subjects : []model.Subject {},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
