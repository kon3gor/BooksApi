package fetchers

import (
	"log"
	"dev/kon3gor/booksapi/mapper"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/openlib"
	"encoding/json"
	"fmt"
	"strconv"
)

func SubjectFetcher(subjectId string, limit int, ch chan model.Subject) { 
	path := fmt.Sprintf("subjects/%s", subjectId)
	params := make(map[string]string, 0)
	if limit > 0  {
		params["limit"] = strconv.Itoa(limit)
	}
	result := openlib.FetchSmthFromOpenLib("GET", path, params)

	defer result.Body.Close()
	
	var resultModel model.MainSubject
	if err := json.NewDecoder(result.Body).Decode(&resultModel); err != nil {
		log.Fatalln(err)
	}

	subject, err := mapper.MapFromMainSubject(resultModel)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- subject

}
