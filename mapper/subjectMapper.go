package mapper

import (
	"log"
	"dev/kon3gor/booksapi/model"
)

func MapFromMainSubject(mainSubject model.MainSubject) (model.Subject, error) {
	books := make([]model.Book, len(mainSubject.Works))
	for i, item := range mainSubject.Works {
		book, err := MapFromSubjectBook(item)
		if err != nil {
			log.Fatalln(err)
		}
		books[i] = book
	}
	return model.Subject {
		Name : mainSubject.Name,
		Books : books,
	}, nil
}
