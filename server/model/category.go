package model

type Category struct {
	Id    int
	Title string
}

type AddCategory struct {
	Title string
}

type UpdateCategory struct {
	Id    int
	Title *string
}
