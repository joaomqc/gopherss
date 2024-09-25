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
	IsMuted     bool
	OriginalId  string
	FeedId      int
}

type UpdateEntriesInput struct {
	Ids     []int
	Read    *bool
	Starred *bool
	Muted   *bool
}

type UpdateEntryInput struct {
	Read    *bool
	Starred *bool
	Muted   *bool
}

type ListEntriesInput struct {
	BaseQuery
	Category *int    `form:"category"`
	Feed     *int    `form:"feed"`
	Starred  *bool   `form:"starred"`
	Read     *bool   `form:"read"`
	Search   *string `form:"search"`
}

type MarkEntriesInput struct {
	Category *int        `form:"category"`
	Feed     *int        `form:"feed"`
	Before   time.Time   `form:"before"`
	As       EntryStatus `form:"as"`
}

type MarkEntryInput struct {
	As EntryStatus `form:"as"`
}

type EntryStatus string

const (
	ReadEntryStatus   EntryStatus = "read"
	UnreadEntryStatus EntryStatus = "unread"
)
