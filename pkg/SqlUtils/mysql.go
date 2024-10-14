package SqlUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlDatabase struct {
	dsn string
	db  *sql.DB
}

func NewMysqlDatabase(user, password, address, port string) (*MysqlDatabase, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, address, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MysqlDatabase{dsn: dsn, db: db}, nil
}

func (m *MysqlDatabase) GetSourceDatabases() ([]string, error) {
	rows, err := m.db.Query("SHOW DATABASES")

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var databases []string
	for rows.Next() {
		var database string
		if err = rows.Scan(&database); err != nil {
			return nil, err
		}
		databases = append(databases, database)
	}
	return databases, nil
}

func (m *MysqlDatabase) GetSourceTables(database string) ([]string, error) {
	rows, err := m.db.Query(fmt.Sprintf("SHOW TABLES FROM %s", database))
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	if err != nil {
		return nil, err
	}
	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, err
}

func (m *MysqlDatabase) GetSourceTableColumns(database string, table string) ([]Column, error) {
	sqlSentence := fmt.Sprintf("SHOW COLUMNS FROM %s.%s;", database, table)

	rows, err := m.db.Query(sqlSentence)
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
		if err = rows.Scan(&column.Field, &column.Type, &column.Null, &column.Key, &column.Default, &column.Extra); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}

type IColumn struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (m *MysqlDatabase) ExecuteSql(sqlSentence string) (map[string]interface{}, error) {

	rows, err := m.db.Query(sqlSentence)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error getting columns: %v", err)
	}

	fmt.Println("columns:====", columns)

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("error getting column types: %v", err)
	}
	fmt.Println("columnTypes:====", columnTypes)

	// Prepare the slice for the rows
	var results []map[string]interface{}

	// Scan the rows
	for rows.Next() {
		// Create a slice to hold column values
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(sql.RawBytes)
		}

		if err := rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Create a map to hold the row data
		rowData := make(map[string]interface{})
		for i, col := range columns {
			value := values[i].(*sql.RawBytes)
			rowData[col] = string(*value)
		}

		results = append(results, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	var columnMetadata []map[string]string
	for _, columnType := range columnTypes {
		columnInfo := map[string]string{
			"name": columnType.Name(),
			"type": columnType.DatabaseTypeName(),
		}
		columnMetadata = append(columnMetadata, columnInfo)
	}

	ret := make(map[string]interface{}, 0)
	ret["columns"] = columnMetadata
	ret["rows"] = results

	return ret, nil
}
