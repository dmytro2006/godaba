package db

type Explorer interface {
	Connect(connection string) error
	ListTables() ([]string, error)
	GetTableSchema(table string) ([]string, error)
	GetTableRows(table string, limit, offset int) ([][]string, error)
	ExecuteQuery(query string) ([]string, [][]string, error)
}

func NewExplorer(dbType, connection string) (Explorer, error) {
	var exp Explorer
	switch dbType {
	case "postgres":
		exp = &Postgres{}
	case "sqlite":
		exp = &SQLite{}
	default:
		return nil, UnknownDriverError{"Unknown DB driver: " + dbType}
	}
	err := exp.Connect(connection)
	if err != nil {
		return nil, err
	}
	return exp, nil
}
