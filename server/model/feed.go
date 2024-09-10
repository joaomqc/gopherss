package model

import "net/url"

type Feed struct {
	Id         int
	Title      string
	FeedUrl    url.URL
	WebsiteUrl url.URL
	CategoryId int
	Visibility Visibility
}

type AddFeed struct {
	Title      string
	FeedUrl    url.URL
	WebsiteUrl url.URL
	CategoryId int
	Visibility Visibility
}

type UpdateFeed struct {
	Title      *string
	FeedUrl    *url.URL
	WebsiteUrl *url.URL
	CategoryId *int
	Visibility *Visibility
}
