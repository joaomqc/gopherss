package db

import (
	"errors"
	"fmt"
	"gopherss/model"
)

type EntriesRepository struct{}

func (r *EntriesRepository) UpdateMany(input model.UpdateEntries) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	setClauses := []setClause{}

	if input.Read != nil {
		setClauses = append(setClauses, setClause{
			field: "is_read",
			value: *input.Read,
		})
	}

	if input.Starred != nil {
		setClauses = append(setClauses, setClause{
			field: "is_starred",
			value: *input.Read,
		})
	}

	whereClauses := []whereClause{
		{
			field: "id",
			op:    "in",
			value: input.Ids,
		},
	}

	q, args := buildUpdateQuery(updateQuery{
		table:        "entries",
		whereClauses: whereClauses,
		setClauses:   setClauses,
	})

	_, err = db.Exec(q, args...)
	if err != nil {
		return fmt.Errorf("entries UpdateMany: %w", err)
	}

	return nil
}

func (r *EntriesRepository) GetMany(query model.EntryListQuery) ([]model.Entry, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	entries := []model.Entry{}

	orderBy := "published_on"
	if query.Order != nil {
		orderBy = *query.Order
	}

	sort := "DESC"
	if query.Sort != nil {
		sort = string(*query.Sort)
	}

	limit := 100
	if query.Limit != nil {
		limit = *query.Limit
	}

	whereClauses := []whereClause{}

	if query.Category != nil && query.Feed != nil {
		return nil, errors.New("cannot filter entries by both category and feed")
	}

	if query.Category != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "category",
			op:    "=",
			value: *query.Category,
		})
	}

	if query.Feed != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "feed",
			op:    "=",
			value: *query.Feed,
		})
	}

	if query.Read != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "is_read",
			op:    "=",
			value: *query.Read,
		})
	}

	if query.Starred != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "is_starred",
			op:    "=",
			value: *query.Starred,
		})
	}

	if query.Search != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "title",
			op:    "like",
			value: *query.Search,
		})
	}

	q, args := buildSelectQuery(selectQuery{
		fields:       []string{"id", "title", "content", "link", "author", "published_on", "collected_on", "is_read", "category", "original_id", "feed_id"},
		table:        "entries",
		whereClauses: whereClauses,
		orderBy:      orderBy,
		sort:         sort,
		limit:        limit,
		offset:       query.Offset,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("entries GetMany: %w", err)
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
			&entry.IsStarred,
			&entry.Category,
			&entry.OriginalId,
			&entry.FeedId)

		if err != nil {
			return nil, fmt.Errorf("entries GetMany: %w", err)
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("entries GetMany: %w", err)
	}
	return entries, nil
}
