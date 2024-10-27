package svc

import (
	"errors"
	"gopherss/db"
	"gopherss/model"
	"net/url"
)

type FeedsService struct{}

var feedsRepository = db.FeedsRepository{}
var rssService = RssService{}

func (s FeedsService) Get(id int) (*model.Feed, error) {
	feed, err := feedsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (s FeedsService) GetIcon(id int) (string, error) {
	return "", errors.New("not implemented")
}

func (s FeedsService) GetMany(input model.ListFeedsInput) ([]model.Feed, error) {
	feeds, err := feedsRepository.GetMany(input)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func (s FeedsService) Create(input model.AddFeedInput) (int, error) {
	feed := model.Feed{
		FeedUrl:    input.FeedUrl,
		CategoryId: input.CategoryId,
	}
	parsedFeed, err := rssService.ParseFeed(feed)
	if err != nil {
		return 0, err
	}
	websiteUrl, err := url.Parse(parsedFeed.Link)
	if err != nil {
		return 0, err
	}
	feed.WebsiteUrl = *websiteUrl
	feed.Title = parsedFeed.Title
	feedId, err := feedsRepository.Insert(feed)
	if err != nil {
		return 0, err
	}
	feed.Id = feedId
	entries, err := rssService.GetEntries(feed, *parsedFeed)
	if err != nil {
		return 0, err
	}
	err = entriesRepository.InsertMany(entries)
	if err != nil {
		return 0, err
	}
	return feedId, nil
}

func (s FeedsService) Update() error {
	return errors.New("not implemented")
}

func (s FeedsService) Delete(id int) error {
	err := feedsRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s FeedsService) Refresh(id int) error {
	return errors.New("not implemented")
}

func (s FeedsService) RefreshMany() error {
	return errors.New("not implemented")
}
