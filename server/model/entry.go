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
	Category    string
	OriginalId  string
	FeedId      int
}

type UpdateEntries struct {
	Ids    []int
	Status Status
}
