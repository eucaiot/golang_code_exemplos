package domain

type PostById struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Reactions `json:"reactions"`
	Views     int `json:"views"`
	UserID    int `json:"userId"`
}

type Reactions struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type PostsSearch struct {
	Posts []PostById
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}
