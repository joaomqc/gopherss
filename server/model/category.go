package model

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type AddCategoryInput struct {
	Title string `json:"title"`
}

type UpdateCategoryInput struct {
	Title *string `json:"title,omitempty"`
}
