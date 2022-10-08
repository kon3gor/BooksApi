package fetchers

import(
	"dev/kon3gor/booksapi/openlib"
	"encoding/json"
	"dev/kon3gor/booksapi/model"
	"dev/kon3gor/booksapi/mapper"
)

func TrendingFetcher() ([]model.Book, error){
	result := openlib.FetchSmthFromOpenLib("GET", "trending/now", nil)

	defer result.Body.Close()
	
	var resultModel model.TrendingNow
	if err := json.NewDecoder(result.Body).Decode(&resultModel); err != nil {
		return []model.Book{}, err
	}
	
	books, err := mapper.MapFromTrendingNow(resultModel)
	if err != nil {
		return []model.Book{}, err 
	}
	
	return books, nil
}
