package SqlUtils

import (
	"database/sql"
	"fmt"
)

func GetDatabaseInstance(dbType, user, password, address, port string) (IDatabase, error) {
	switch dbType {
	case "clickhouse":
		return NewClickHouseDatabase(user, password, address, port)
	case "mysql":
		return NewMysqlDatabase(user, password, address, port)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

type Column struct {
	Field   string         `json:"name"`
	Type    string         `json:"type"`
	Null    string         `json:"-"`
	Key     string         `json:"-"`
	Default sql.NullString `json:"-"`
	Extra   string         `json:"-"`
}
