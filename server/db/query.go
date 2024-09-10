package db

import (
	"fmt"
	"strings"
)

type selectQuery struct {
	fields  []string
	from    string
	where   []whereClause
	orderBy string
	sort    string
	limit   int
	offset  *int
}

type whereClause struct {
	field string
	op    string
	value any
}

func buildSelectQuery(query selectQuery) (string, []any, error) {
	args := []any{}
	q := fmt.Sprint("SELECT ", strings.Join(query.fields, ", "), " FROM ", query.from)

	if len(query.where) > 0 {
		q += " WHERE 1=1 "
	}
	for _, where := range query.where {
		q += fmt.Sprintf(" AND %s %s ?", where.field, where.op)
		args = append(args, where.value)
	}

	q += " ORDER BY ? ? "
	args = append(args, query.orderBy, query.sort)

	q += " LIMIT ? "
	args = append(args, query.limit)

	if query.offset != nil {
		q += " OFFSET ? "
		args = append(args, *query.offset)
	}

	return q, args, nil
}
