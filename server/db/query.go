package db

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type updateQuery struct {
	setClauses   []setClause
	table        string
	whereClauses []whereClause
}

type setClause struct {
	field string
	value any
}

type selectQuery struct {
	fields       []string
	table        string
	whereClauses []whereClause
	joinClauses  []joinClause
	orderBy      *string
	sort         *string
	limit        *int
	offset       *int
}

type whereClause struct {
	field string
	op    string
	value any
}

type joinClause struct {
	table string
	on    string
}

func buildSelectQuery(query selectQuery) (string, []any) {
	args := []any{}
	q := fmt.Sprint("SELECT ", strings.Join(query.fields, ", "), " FROM ", query.table)

	if len(query.whereClauses) > 0 {
		q += " WHERE "
		wheres := []string{}
		for _, where := range query.whereClauses {
			wheres = append(wheres, fmt.Sprintf("%s %s ?", where.field, where.op))
			args = append(args, where.value)
		}
		q += strings.Join(wheres, " AND ")
	}

	if query.sort != nil {
		q += fmt.Sprint(" ORDER BY ? ", *query.sort)
		args = append(args, *query.orderBy)
	}

	if query.limit != nil {
		q += " LIMIT ? "
		args = append(args, *query.limit)
	}

	if query.offset != nil {
		q += " OFFSET ? "
		args = append(args, *query.offset)
	}

	return q, args
}

func buildUpdateQuery(query updateQuery) (string, []any) {
	args := []any{}
	q := fmt.Sprint("UPDATE ", query.table, " SET ")

	sets := []string{}
	for _, set := range query.setClauses {
		sets = append(sets, fmt.Sprint(set.field, " = ?"))
		args = append(args, set.value)
	}
	q += strings.Join(sets, ", ")

	wheres := []string{}
	for _, where := range query.whereClauses {
		if where.op == "in" {
			wheres = append(wheres, fmt.Sprintf(" AND %s IN (?)", where.field))
			args = append(args, pq.Array(where.value))
		} else {
			wheres = append(wheres, fmt.Sprintf(" AND %s %s ?", where.field, where.op))
			args = append(args, where.value)
		}
	}
	q += strings.Join(wheres, " AND ")

	return q, args
}
