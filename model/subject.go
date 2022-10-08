package model

type Subject struct {
	Name string `json:"name"`
	Books []Book `json:"books"`
}

type MainSubject struct {
	Name string `json:"name"`
	Works []SubjectBook `json:"works"`
}
