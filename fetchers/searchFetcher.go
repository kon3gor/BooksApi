package fetchers

import (
	"log"
	"dev/kon3gor/booksapi/mapper"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/openlib"
	"encoding/json"
	"strconv"
)

func SearchFetcher(text string, limit int, ch chan []model.Book) {
	params := map[string]string {
		"q" : text,
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	result := openlib.FetchSmthFromOpenLib("GET", "search", params)
	defer result.Body.Close()

	var resultModel model.SearchResult
	if err := json.NewDecoder(result.Body).Decode(&resultModel); err != nil {
		log.Fatalln(err)
	}

	books, err := mapper.MapFromSearchResult(resultModel)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- books
}
