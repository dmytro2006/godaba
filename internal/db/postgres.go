package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect(connection string) error {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) ListTables() ([]string, error) {
	rows, err := p.db.Query(`
        SELECT tablename
        FROM pg_catalog.pg_tables
        WHERE schemaname = 'public';
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tables = append(tables, name)
	}

	return tables, nil
}

func (p *Postgres) ExecuteQuery(query string) ([]string, [][]string, error) {
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	var result [][]string
	for rows.Next() {
		row := make([]string, len(columns))
		ptrs := make([]any, len(columns))
		for i := range columns {
			ptrs[i] = &row[i]
		}

		err := rows.Scan(ptrs...)
		if err != nil {
			return nil, nil, err
		}
		result = append(result, row)
	}

	return columns, result, nil

}

func (p *Postgres) GetTableRows(table string, limit, offset int) ([][]string, error) {
	query := fmt.Sprintf(
		`
		SELECT *
		FROM %s
		LIMIT %d OFFSET %d
		`,
		table,
		limit,
		offset,
	)
	_, rows, err := p.ExecuteQuery(query)
	return rows, err
}

func (p *Postgres) GetTableSchema(table string) ([]string, error) {
	query := fmt.Sprintf(
		`
		SELECT *
		FROM %s
		LIMIT 0 OFFSET 0
		`,
		table,
	)
	columns, _, err := p.ExecuteQuery(query)
	return columns, err
}
