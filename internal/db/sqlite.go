package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func (s *SQLite) Connect(connection string) error {
	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *SQLite) ListTables() ([]string, error) {
	rows, err := s.db.Query(`
        SELECT name 
		FROM sqlite_schema 
		WHERE type = 'table' 
		AND name NOT LIKE 'sqlite_%'
		AND name NOT IN ('__EFMigrationsHistory', 'InnaTabelaNarzędziowa')
		ORDER BY name;
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

func (s *SQLite) ExecuteQuery(query string) ([]string, [][]string, error) {
	rows, err := s.db.Query(query)
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

func (s *SQLite) GetTableRows(table string, limit, offset int) ([][]string, error) {
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
	_, rows, err := s.ExecuteQuery(query)
	return rows, err
}

func (s *SQLite) GetTableSchema(table string) ([]string, error) {
	query := fmt.Sprintf(
		`
		SELECT *
		FROM %s
		LIMIT 0 OFFSET 0
		`,
		table,
	)
	columns, _, err := s.ExecuteQuery(query)
	return columns, err
}
