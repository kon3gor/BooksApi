package model

type AuthorMain struct {
	Author KeyValue `json:"author"`
}

type SubjectAuthor struct {
	Name string `json:"name"`
}
