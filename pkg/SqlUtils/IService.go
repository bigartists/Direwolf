package SqlUtils

type IDatabase interface {
	GetSourceDatabases() ([]string, error)
	GetSourceTables(database string) ([]string, error)
	GetSourceTableColumns(database string, table string) ([]Column, error)
	ExecuteSql(sql string) (map[string]interface{}, error)
}
