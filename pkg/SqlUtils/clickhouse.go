package SqlUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"net/url"
)

type ClickHouseDatabase struct {
	dsn string
	db  *sql.DB
}

func NewClickHouseDatabase(user, password, address, port string) (*ClickHouseDatabase, error) {
	encodedPassword := url.QueryEscape(password)
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/", user, encodedPassword, address, port)
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}
	return &ClickHouseDatabase{dsn: dsn, db: db}, nil
}

func (c *ClickHouseDatabase) GetSourceDatabases() ([]string, error) {
	rows, err := c.db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var database string
		if err := rows.Scan(&database); err != nil {
			return nil, err
		}
		databases = append(databases, database)
	}

	return databases, nil
}

func (c *ClickHouseDatabase) GetSourceTables(database string) ([]string, error) {
	rows, err := c.db.Query(fmt.Sprintf("SHOW TABLES IN %s", database))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func (c *ClickHouseDatabase) GetSourceTableColumns(database, table string) ([]Column, error) {

	sqlSentence := fmt.Sprintf("SHOW COLUMNS FROM %s.%s;", database, table)

	rows, err := c.db.Query(sqlSentence)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var columns []Column
	for rows.Next() {
		//var field, fieldType, null, key, extra string
		//var defaultValue sql.NullString // beacuse  name "defaultValue": converting NULL to string is unsupported Null 不支持强转为 string
		var column Column
		if err = rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}

// todo clickhouse impl 未完成
func (c *ClickHouseDatabase) ExecuteSql(sql string) (map[string]interface{}, error) {
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := [][]string{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		result := make([]string, len(columns))
		for i, value := range values {
			if value != nil {
				result[i] = fmt.Sprintf("%s", value)
			} else {
				result[i] = "NULL"
			}
		}
		results = append(results, result)
	}
	panic("impl it,")
	return nil, nil
}
