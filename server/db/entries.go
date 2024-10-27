package db

import (
	"errors"
	"fmt"
	"gopherss/model"
	"strings"
)

type EntriesRepository struct{}

func (r *EntriesRepository) GetMany(input model.ListEntriesInput) ([]model.Entry, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	entries := []model.Entry{}

	orderBy := "published_on"
	if input.Order != nil {
		orderBy = *input.Order
	}

	sort := "DESC"
	if input.Sort != nil {
		switch *input.Sort {
		case model.AscendingSort:
			sort = "ASC"
		case model.DescendingSort:
			sort = "DESC"
		}
	}

	limit := 100
	if input.Limit != nil {
		limit = *input.Limit
	}

	whereClauses := []whereClause{}

	if input.Category != nil && input.Feed != nil {
		return nil, errors.New("cannot filter entries by both category and feed")
	}

	if input.Category != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "category_id",
			op:    "=",
			value: *input.Category,
		})
	}

	if input.Feed != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "feed",
			op:    "=",
			value: *input.Feed,
		})
	}

	if input.Read != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "is_read",
			op:    "=",
			value: *input.Read,
		})
	}

	if input.Starred != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "is_starred",
			op:    "=",
			value: *input.Starred,
		})
	}

	if input.Search != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "title",
			op:    "like",
			value: *input.Search,
		})
	}

	q, args := buildSelectQuery(selectQuery{
		fields:       []string{"id", "title", "content", "link", "author", "published_on", "collected_on", "is_read", "is_starred", "original_id", "feed_id"},
		table:        "entries",
		whereClauses: whereClauses,
		joinClauses: []joinClause{{
			table: "feeds",
			on:    "entries.feed_id = feeds.id",
		}},
		orderBy: &orderBy,
		sort:    &sort,
		limit:   &limit,
		offset:  input.Offset,
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

func (r *EntriesRepository) UpdateMany(input model.UpdateEntriesInput) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	setClauses := []setClause{}

	if input.IsRead != nil {
		setClauses = append(setClauses, setClause{
			field: "is_read",
			value: *input.IsRead,
		})
	}

	if input.IsStarred != nil {
		setClauses = append(setClauses, setClause{
			field: "is_starred",
			value: *input.IsStarred,
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

func (r *EntriesRepository) Update(id int, input model.UpdateEntryInput) (bool, error) {
	db, err := GetDb()
	if err != nil {
		return false, err
	}

	setClauses := []setClause{}

	if input.IsRead != nil {
		setClauses = append(setClauses, setClause{
			field: "is_read",
			value: *input.IsRead,
		})
	}

	if input.IsStarred != nil {
		setClauses = append(setClauses, setClause{
			field: "is_starred",
			value: *input.IsStarred,
		})
	}

	whereClauses := []whereClause{
		{
			field: "id",
			op:    "=",
			value: id,
		},
	}

	q, args := buildUpdateQuery(updateQuery{
		table:        "entries",
		whereClauses: whereClauses,
		setClauses:   setClauses,
	})

	res, err := db.Exec(q, args...)
	if err != nil {
		return false, fmt.Errorf("entries Update: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("entries Update: %w", err)
	}

	return rows > 0, nil
}

func (r *EntriesRepository) MarkMany(input model.MarkEntriesInput) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	var isRead bool
	switch input.As {
	case model.ReadEntryStatus:
		isRead = true
	case model.UnreadEntryStatus:
		isRead = false
	default:
		return errors.New("unknown status")
	}

	setClauses := []setClause{
		{
			field: "is_read",
			value: isRead,
		},
	}

	whereClauses := []whereClause{
		{
			field: "collected_on",
			op:    "<",
			value: input.Before,
		},
	}

	if input.Category != nil && input.Feed != nil {
		return errors.New("cannot filter entries by both category and feed")
	}

	if input.Category != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "category",
			op:    "=",
			value: *input.Category,
		})
	}

	if input.Feed != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "feed",
			op:    "=",
			value: *input.Feed,
		})
	}

	q, args := buildUpdateQuery(updateQuery{
		table:        "entries",
		whereClauses: whereClauses,
		setClauses:   setClauses,
	})

	_, err = db.Exec(q, args...)
	if err != nil {
		return fmt.Errorf("entries MarkMany: %w", err)
	}

	return nil
}

func (r *EntriesRepository) Mark(id int, input model.MarkEntryInput) (bool, error) {
	db, err := GetDb()
	if err != nil {
		return false, err
	}

	var isRead bool
	switch input.As {
	case model.ReadEntryStatus:
		isRead = true
	case model.UnreadEntryStatus:
		isRead = false
	default:
		return false, errors.New("unknown status")
	}

	setClauses := []setClause{
		{
			field: "is_read",
			value: isRead,
		},
	}

	whereClauses := []whereClause{
		{
			field: "id",
			op:    "=",
			value: id,
		},
	}

	q, args := buildUpdateQuery(updateQuery{
		table:        "entries",
		whereClauses: whereClauses,
		setClauses:   setClauses,
	})

	res, err := db.Exec(q, args...)
	if err != nil {
		return false, fmt.Errorf("entries Mark: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("entries Mark: %w", err)
	}

	return rows > 0, nil
}

func (r *EntriesRepository) Get(id int) (*model.Entry, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}

	whereClauses := []whereClause{{
		field: "id",
		op:    "=",
		value: id,
	}}

	q, args := buildSelectQuery(selectQuery{
		fields:       []string{"id", "title", "content", "link", "author", "published_on", "collected_on", "is_read", "is_starred", "original_id", "feed_id"},
		table:        "entries",
		whereClauses: whereClauses,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("entries Get: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("entries Get: %w", err)
		}
		return nil, nil
	}

	var entry model.Entry
	err = rows.Scan(
		&entry.Id,
		&entry.Title,
		&entry.Content,
		&entry.Link,
		&entry.Author,
		&entry.PublishedOn,
		&entry.CollectedOn,
		&entry.IsRead,
		&entry.IsStarred,
		&entry.OriginalId,
		&entry.FeedId)

	if err != nil {
		return nil, fmt.Errorf("entries Get: %w", err)
	}

	return &entry, nil
}

func (r *EntriesRepository) InsertMany(entries []model.Entry) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	inserts := []string{}
	vals := []any{}

	for _, entry := range entries {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?, ?)")
		vals = append(vals, entry.Title, entry.Content, entry.Link, entry.Author, entry.PublishedOn, entry.CollectedOn, entry.OriginalId, entry.FeedId)
	}

	query := fmt.Sprintf("INSERT INTO entries(title, content, link, author, published_on, colected_on, original_id, feed_id) VALUES %s", strings.Join(inserts, ", "))

	_, err = db.Exec(query, vals...)
	if err != nil {
		return fmt.Errorf("entries InsertMany: %w", err)
	}

	return err
}
