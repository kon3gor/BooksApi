package model

type Book struct {
	Key string `json:"key"`
	Subjects []string `json:"subjects"`
	Title string `json:"title"`
	Description string `json:"description"`
	Image string `json:"image,omitempty"`
	Author string `json:"author,omitempty"`
}

type MainBook struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Subjects []string `json:"subjects"`
	Authors []AuthorMain `json:"authors"`
	Covers []int `json:"covers"`
	Description string `json:"description"`
}

type SubjectBook struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Subjects []string `json:"subject"`
	Authors []SubjectAuthor `json:"authors"`
	Cover int `json:"cover_id"`
	Description string `json:"description"`
}

type SearchBook struct {
	Key string `json:"key"`
	Title string `json:"title"`
	Author []string `json:"author_name"`
	Cover int `json:"cover_i"`
}
