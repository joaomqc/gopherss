package db

import (
	"fmt"
	"gopherss/model"
)

type ArticlesRepository struct{}

func (r *ArticlesRepository) GetAll() ([]model.Article, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	articles := []model.Article{}
	rows, err := db.Query(`
		SELECT
			id,
			title,
			content,
			link,
			author,
			published_on,
			collected_on,
			is_read,
			category,
			original_id,
			feed_id
		FROM articles
	`)
	if err != nil {
		return nil, fmt.Errorf("articles getAll: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var article model.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Content,
			&article.Link,
			&article.Author,
			&article.PublishedOn,
			&article.CollectedOn,
			&article.IsRead,
			&article.Category,
			&article.OriginalId,
			&article.FeedId)

		if err != nil {
			return nil, fmt.Errorf("articles getAll: %w", err)
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("articles getAll: %w", err)
	}
	return articles, nil
}
