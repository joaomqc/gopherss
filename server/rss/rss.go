package rss

import (
	"fmt"
	"gopherss/model"
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
)

func parseFeed(feedUrl url.URL) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	// TODO support auth
	feed, err := fp.ParseURL(feedUrl.String())
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func toEntries(feed model.Feed, items []*gofeed.Item) ([]model.Entry, error) {
	entries := []model.Entry{}
	for _, item := range items {
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
			IsMuted:     false,
			OriginalId:  item.GUID,
			FeedId:      feed.Id,
		})
	}

	return entries, nil
}

func refreshFeed(feed model.Feed) error {
	parsedFeed, err := parseFeed(feed.FeedUrl)
	if err != nil {
		return nil
	}

	entries, err := toEntries(feed, parsedFeed.Items)
	if err != nil {
		return err
	}
	//TODO do something with the entries
	fmt.Println(entries)
	return nil
}
