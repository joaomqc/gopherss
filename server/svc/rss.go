package svc

import (
	"gopherss/model"
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
)

type RssService struct{}

func (s RssService) ParseFeed(feed model.Feed) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	// TODO support auth
	parsedFeed, err := fp.ParseURL(feed.FeedUrl.String())
	if err != nil {
		return nil, err
	}

	return parsedFeed, nil
}

func (s RssService) GetEntries(feed model.Feed, parsedFeed gofeed.Feed) ([]model.Entry, error) {
	entries := []model.Entry{}
	for _, item := range parsedFeed.Items {
		link, err := url.Parse(item.Link)
		if err != nil {
			return nil, err
		}
		entries = append(entries, model.Entry{
			Title:       item.Title,
			Content:     item.Content,
			Link:        *link,
			Author:      item.Author.Name,
			PublishedOn: *item.PublishedParsed,
			CollectedOn: time.Now(),
			IsRead:      false,
			IsStarred:   false,
			OriginalId:  item.GUID,
			FeedId:      feed.Id,
		})
	}
	return entries, nil
}
