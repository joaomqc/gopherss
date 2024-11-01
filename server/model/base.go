package model

type BaseQuery struct {
	Offset *int      `form:"offset"`
	Limit  *int      `form:"limit"`
	Order  *string   `form:"order"`
	Sort   *SortType `form:"sort"`
	Search *string   `form:"search"`
}

type SortType string

const (
	AscendingSort  SortType = "ascending"
	DescendingSort SortType = "descending"
)
