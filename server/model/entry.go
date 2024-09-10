package model

import (
	"net/url"
	"time"
)

type Entry struct {
	Id          int
	Title       string
	Content     string
	Link        url.URL
	Author      string
	PublishedOn time.Time
	CollectedOn time.Time
	IsRead      bool
	IsStarred   bool
	Category    string
	OriginalId  string
	FeedId      int
}

type UpdateEntries struct {
	Ids    []int
	Status Status
}

type UpdateEntry struct {
	Title       *string
	Content     *string
	Link        *url.URL
	Author      *string
	PublishedOn *time.Time
	CollectedOn *time.Time
	IsRead      *bool
	Category    *string
	OriginalId  *string
	FeedId      *int
}

type EntryListQuery struct {
	BaseQuery
	Category *int    `form:"category"`
	Feed     *int    `form:"feed"`
	Starred  *bool   `form:"starred"`
	Read     *bool   `form:"read"`
	Search   *string `form:"search"`
}
