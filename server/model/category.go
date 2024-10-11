package model

type Category struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Visibility int    `json:"visibility"`
}

type AddCategoryInput struct {
	Title      string `json:"title"`
	Visibility int    `json:"visibility"`
}

type UpdateCategoryInput struct {
	Title      *string `json:"title,omitempty"`
	Visibility *int    `json:"visibility,omitempty"`
}

type ListCategoriesInput struct {
	BaseQuery
	ShowHidden bool `form:"showHidden"`
}

type CategoryVisibility int

const (
	ShowCategoryVisibility      CategoryVisibility = 1
	DoNotShowCategoryVisibility CategoryVisibility = 2
)
