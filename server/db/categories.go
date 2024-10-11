package db

import (
	"fmt"
	"gopherss/model"
)

type CategoriesRepository struct{}

func (r *CategoriesRepository) GetMany(input model.ListCategoriesInput) ([]model.Category, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	categories := []model.Category{}

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

	if input.Search != nil {
		whereClauses = append(whereClauses, whereClause{
			field: "title",
			op:    "like",
			value: "%" + *input.Search + "%",
		})
	}

	visibilities := []model.CategoryVisibility{model.ShowCategoryVisibility}
	if input.ShowHidden {
		visibilities = append(visibilities, model.DoNotShowCategoryVisibility)
	}
	whereClauses = append(whereClauses, whereClause{
		field: "visibility",
		op:    "in",
		value: visibilities,
	})

	q, args := buildSelectQuery(selectQuery{
		fields:       []string{"id", "title"},
		table:        "categories",
		whereClauses: whereClauses,
		orderBy:      &orderBy,
		sort:         &sort,
		limit:        &limit,
		offset:       input.Offset,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("categories GetMany: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.Id,
			&category.Title)

		if err != nil {
			return nil, fmt.Errorf("categories GetMany: %w", err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("categories GetMany: %w", err)
	}

	return categories, nil
}

func (r *CategoriesRepository) Get(id int) (*model.Category, error) {
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
		fields:       []string{"id", "title"},
		table:        "categories",
		whereClauses: whereClauses,
	})

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("categories Get: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("categories Get: %w", err)
		}
		return nil, nil
	}

	var category model.Category
	err = rows.Scan(
		&category.Id,
		&category.Title)

	if err != nil {
		return nil, fmt.Errorf("categories Get: %w", err)
	}

	return &category, nil
}

func (r *CategoriesRepository) Update(id int, input model.UpdateCategoryInput) (bool, error) {
	db, err := GetDb()
	if err != nil {
		return false, err
	}

	setClauses := []setClause{}

	if input.Title != nil {
		setClauses = append(setClauses, setClause{
			field: "title",
			value: *input.Title,
		})
	}

	if input.Visibility != nil {
		setClauses = append(setClauses, setClause{
			field: "visibility",
			value: *input.Visibility,
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
		table:        "categories",
		whereClauses: whereClauses,
		setClauses:   setClauses,
	})

	res, err := db.Exec(q, args...)
	if err != nil {
		return false, fmt.Errorf("categories Update: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("categories Update: %w", err)
	}

	return rows > 0, nil
}

func (r *CategoriesRepository) Insert(input model.AddCategoryInput) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	query := `INSERT INTO categories(title, visibility) VALUES (?, ?)`

	_, err = db.Exec(query, input.Title, input.Visibility)
	if err != nil {
		return fmt.Errorf("categories Insert: %w", err)
	}

	return err
}

func (r *CategoriesRepository) Delete(id int, keepFeeds bool) error {
	db, err := GetDb()
	if err != nil {
		return err
	}

	query := `DELETE FROM feeds WHERE category_id = ?; DELETE FROM categories WHERE id = ?`
	if keepFeeds {
		query = `UPDATE feeds SET category_id = 1 WHERE category_id = ?; DELETE FROM categories WHERE id = ?`
	}

	_, err = db.Exec(query, id, id)
	if err != nil {
		return fmt.Errorf("categories Delete: %w", err)
	}

	return err
}
