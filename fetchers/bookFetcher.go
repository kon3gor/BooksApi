package fetchers

import(
	"log"
	"fmt"
	"dev/kon3gor/booksapi/openlib"
	"encoding/json"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/mapper"
)

func BookFetcher(bookId string, ch chan model.Book) {
	path := fmt.Sprintf("works/%s", bookId)
	result := openlib.FetchSmthFromOpenLib("GET", path, nil)

	defer result.Body.Close()
	
	var resultModel model.MainBook
	if err := json.NewDecoder(result.Body).Decode(&resultModel); err != nil {
		log.Fatalln(err)
	}

	book, err := mapper.MapFromMainBook(resultModel)
	if err != nil {
		log.Fatalln(err)
	}
	
	ch <- book
}
