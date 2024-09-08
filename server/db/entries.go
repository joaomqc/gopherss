package db

import (
	"fmt"
	"gopherss/model"
)

type EntriesRepository struct{}

func (r *EntriesRepository) GetAll() ([]model.Entry, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	entries := []model.Entry{}
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
		FROM entries
	`)
	if err != nil {
		return nil, fmt.Errorf("entries getAll: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry model.Entry
		err := rows.Scan(
			&entry.Id,
			&entry.Title,
			&entry.Content,
			&entry.Link,
			&entry.Author,
			&entry.PublishedOn,
			&entry.CollectedOn,
			&entry.IsRead,
			&entry.Category,
			&entry.OriginalId,
			&entry.FeedId)

		if err != nil {
			return nil, fmt.Errorf("entries getAll: %w", err)
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("entries getAll: %w", err)
	}
	return entries, nil
}
