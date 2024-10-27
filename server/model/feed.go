package model

import "net/url"

type Feed struct {
	Id         int            `json:"id"`
	Title      string         `json:"title"`
	FeedUrl    url.URL        `json:"feedUrl"`
	WebsiteUrl url.URL        `json:"websiteUrl"`
	CategoryId int            `json:"categoryId"`
	Visibility FeedVisibility `json:"visibility"`
	IsMuted    bool           `json:"isMuted"`
}

type AddFeedInput struct {
	FeedUrl    url.URL `json:"feedUrl"`
	CategoryId int     `json:"categoryId"`
}

type UpdateFeedInput struct {
	Title      *string         `json:"title,omitempty"`
	FeedUrl    *url.URL        `json:"feedUrl,omitempty"`
	CategoryId *int            `json:"categoryId,omitempty"`
	Visibility *FeedVisibility `json:"visibility,omitempty"`
	IsMuted    *bool           `json:"isMuted,omitempty"`
}

type ListFeedsInput struct {
	BaseQuery
	Category   *int `form:"category"`
	ShowHidden bool `form:"showHidden"`
}

type FeedVisibility int

const (
	ShowFeedVisibility           FeedVisibility = 1
	ShowInCategoryFeedVisibility FeedVisibility = 2
	DoNotShowFeedVisibility      FeedVisibility = 3
)
