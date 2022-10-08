package main

import (
	"dev/kon3gor/booksapi/pageapi"
	"dev/kon3gor/booksapi/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	
	//pageapi
	router.GET("/page/lists", pageapi.ListsPageHandler)
	router.GET("/page/profile", pageapi.ProfilePageHandler)
	router.GET("/page/book/:bookId", pageapi.BookPageHandler)
	router.GET("/page/search", pageapi.SearchPageHandler)
	router.GET("/page/search/result", pageapi.SearchResultPageHandler)
	router.GET("/page/subject/:subjectId", pageapi.SubjectPageHandler)

	//handlers
	router.GET("/ping", handlers.PingHandler)
	
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, router))

}
