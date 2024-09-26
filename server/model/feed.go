package model

import "net/url"

type Feed struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	FeedUrl    url.URL    `json:"feedUrl"`
	WebsiteUrl url.URL    `json:"websiteUrl"`
	CategoryId int        `json:"categoryId"`
	Visibility Visibility `json:"visibility"`
}

type AddFeed struct {
	Title      string     `json:"title"`
	FeedUrl    url.URL    `json:"feedUrl"`
	WebsiteUrl url.URL    `json:"websiteUrl"`
	CategoryId int        `json:"categoryId"`
	Visibility Visibility `json:"visibility"`
}

type UpdateFeed struct {
	Title      *string     `json:"title,omitempty"`
	FeedUrl    *url.URL    `json:"feedUrl,omitempty"`
	WebsiteUrl *url.URL    `json:"websiteUrl,omitempty"`
	CategoryId *int        `json:"categoryId,omitempty"`
	Visibility *Visibility `json:"visibility,omitempty"`
}
