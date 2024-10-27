package db

import (
	"fmt"
	"gopherss/model"
)

type FeedsRepository struct{}

func (r *FeedsRepository) GetMany(input model.ListFeedsInput) ([]model.Feed, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	feeds := []model.Feed{}

	orderBy := "id"
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
	visibilities := []model.FeedVisibility{model.ShowFeedVisibility}

	if input.Category != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "category",
			op:    "=",
			value: *input.Category,
		})
		visibilities = append(visibilities, model.ShowInCategoryFeedVisibility)
	}

	if input.Search != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "title",
			op:    "like",
			value: "%" + *input.Search + "%",
		})
	}

	if input.ShowHidden {
		visibilities = append(visibilities, model.DoNotShowFeedVisibility)
	}

	whereClauses = append(whereClauses, whereClause{
		field: "visibility",
		op:    "in",
		value: visibilities,
	})

	q, args := buildSelectQuery(selectQuery{
		fields:       []string{"id", "title", "feedUrl", "websiteUrl", "categoryId", "visibility", "is_muted"},
		table:        "feeds",
		whereClauses: whereClauses,
		orderBy:      &orderBy,
		sort:         &sort,
		limit:        &limit,
		offset:       input.Offset,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("feeds GetMany: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var feed model.Feed
		err := rows.Scan(
			&feed.Id,
			&feed.Title,
			&feed.FeedUrl,
			&feed.WebsiteUrl,
			&feed.CategoryId,
			&feed.Visibility,
			&feed.IsMuted)

		if err != nil {
			return nil, fmt.Errorf("feeds GetMany: %w", err)
		}
		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("feeds GetMany: %w", err)
	}

	return feeds, nil
}

func (r *FeedsRepository) Get(id int) (*model.Feed, error) {
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
		fields:       []string{"id", "title", "feedUrl", "websiteUrl", "categoryId", "visibility", "is_muted"},
		table:        "feeds",
		whereClauses: whereClauses,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("feeds Get: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("feeds Get: %w", err)
		}
		return nil, nil
	}

	var feed model.Feed
	err = rows.Scan(
		&feed.Id,
		&feed.Title,
		&feed.FeedUrl,
		&feed.WebsiteUrl,
		&feed.CategoryId,
		&feed.Visibility,
		&feed.IsMuted)

	if err != nil {
		return nil, fmt.Errorf("feeds Get: %w", err)
	}

	return &feed, nil
}

func (r *FeedsRepository) Insert(input model.Feed) (int, error) {
	db, err := GetDb()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO feeds(title, feedUrl, websiteUrl, categoryId) VALUES (?, ?, ?, ?)`

	var id int
	err = db.QueryRow(query, input.Title, input.FeedUrl, input.WebsiteUrl, input.CategoryId, input.Visibility).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("feeds Insert: %w", err)
	}

	return id, nil
}

func (r *FeedsRepository) Delete(id int) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	query := `DELETE FROM entries WHERE feed_id = ?; DELETE FROM feeds WHERE id = ?`

	_, err = db.Exec(query, id, id)
	if err != nil {
		return fmt.Errorf("feeds Delete: %w", err)
	}

	return err
}
