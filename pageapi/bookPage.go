package pageapi

import (
	"dev/kon3gor/booksapi/fetchers"
	"dev/kon3gor/booksapi/model"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BookPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId := ps.ByName("bookId") 
	ch := make(chan model.Book)

	go fetchers.BookFetcher(bookId, ch)
	result := <-ch

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

