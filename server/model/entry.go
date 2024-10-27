package model

import (
	"net/url"
	"time"
)

type Entry struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Link        url.URL   `json:"link"`
	Author      string    `json:"author"`
	PublishedOn time.Time `json:"publishedOn"`
	CollectedOn time.Time `json:"collectedOn"`
	IsRead      bool      `json:"isRead"`
	IsStarred   bool      `json:"isStarred"`
	OriginalId  string    `json:"originalId"`
	FeedId      int       `json:"feedId"`
}

type UpdateEntriesInput struct {
	Ids       []int `json:"ids"`
	IsRead    *bool `json:"isRead,omitempty"`
	IsStarred *bool `json:"isStarred,omitempty"`
}

type UpdateEntryInput struct {
	IsRead    *bool `json:"isRead,omitempty"`
	IsStarred *bool `json:"isStarred,omitempty"`
}

type ListEntriesInput struct {
	BaseQuery
	Category *int  `form:"category"`
	Feed     *int  `form:"feed"`
	Starred  *bool `form:"starred"`
	Read     *bool `form:"read"`
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
