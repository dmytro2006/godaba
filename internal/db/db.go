package db

type Explorer interface {
	Connect(connection string) error
	ListTables() ([]string, error)
	GetTableSchema(table string) ([]string, error)
	GetTableRows(table string, limit, offset int) ([][]string, error)
	ExecuteQuery(query string) ([]string, [][]string, error)
}

func NewExplorer(dbType, connection string) Explorer {
	switch dbType {
	case "postgres":
		exp := Postgres{}
		exp.Connect(connection)
		return &exp
	}
	return nil
}
