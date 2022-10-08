package mapper

import (
	"dev/kon3gor/booksapi/model"
	"log"
)

func MapFromMainBook(bookMain model.MainBook) (model.Book, error) {
	//todo: check size of a collection here
	author := bookMain.Authors[0].Author.Key
	image := MapLargeImage(bookMain.Covers[0])
	key := MapKey(bookMain.Key)
	book := model.Book {
		Key : key,
		Title : bookMain.Title,
		Subjects : bookMain.Subjects,
		Author : author,
		Description : bookMain.Description,
		Image : image,
	}
	return book, nil
}

func MapFromSubjectBook(subjectBook model.SubjectBook) (model.Book, error) {
	//todo: check size of the collection here
	author := subjectBook.Authors[0].Name
	image := MapSmallImage(subjectBook.Cover)
	key := MapKey(subjectBook.Key)
	book := model.Book {
		Key : key, 
		Title : subjectBook.Title,
		Subjects : subjectBook.Subjects,
		Author : author,
		Description : subjectBook.Description,
		Image : image,
	}
	return book, nil
}

func MapFromSearchBook(trendingBook model.SearchBook) (model.Book, error) {
	//todo: check size of the collection here
	var author string
	if len(trendingBook.Author) > 0 {
		author = trendingBook.Author[0]
	}
	image := MapSmallImage(trendingBook.Cover)
	key := MapKey(trendingBook.Key)
	book := model.Book {
		Key : key, 
		Title : trendingBook.Title,
		Author : author,
		Image : image,
	}
	return book, nil
}

func MapFromTrendingNow(trendingNow model.TrendingNow) ([]model.Book, error) {
	books := mapSearchBooks(trendingNow.Works)
	return books, nil
}

func MapFromSearchResult(searchResult model.SearchResult) ([]model.Book, error) {
	books := mapSearchBooks(searchResult.Docs) 
	return books, nil
}

func mapSearchBooks(searchBooks []model.SearchBook) []model.Book {
	books := make([]model.Book, 0) 
	for _, item := range searchBooks {
		book, err := MapFromSearchBook(item)
		if err != nil {
			log.Fatalln(err)
		} else {
			books = append(books, book)
		}
	}
	return books
}
