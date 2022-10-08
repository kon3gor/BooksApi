package main

import (
	"log"
	"net/http"
	"dev/kon3gor/booksapi/pageapi"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/page/lists", pageapi.ListsPageHandler)
	router.GET("/page/profile", pageapi.ProfilePageHandler)
	router.GET("/page/book/:bookId", pageapi.BookPageHandler)
	router.GET("/page/search", pageapi.SearchPageHandler)
	router.GET("/page/search/result", pageapi.SearchResultPageHandler)
	router.GET("/page/subject/:subjectId", pageapi.SubjectPageHandler)


	log.Fatal(http.ListenAndServe(":8081", router))

}
