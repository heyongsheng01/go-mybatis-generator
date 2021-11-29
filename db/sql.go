package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ccc469/go-mybatis-generator/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

// 初始化数据库连接
func InitDB() {

	var (
		err error
		dsn string
		db  *sql.DB
	)

	config.InitSqlJdbcTypes()
	config.InitFieldTypes()

	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%s?charset=utf8", config.UserName, config.Password, config.Host, config.Port, config.Database)
	if db, err = sql.Open("mysql", dsn); err != nil {
		goto ERROR
	}

	if err = db.Ping(); err != nil {
		goto ERROR
	}

	Db = db
	log.Println(fmt.Sprintf("[%s] init successfully...", config.Database))
	return
ERROR:
	panic(err)
}

// GetTableColumns 查询所有字段
func GetTableColumns(tableName string) (columns []map[string]string) {
	var (
		queryColumns []string
	)
	rows, _ := Db.Query(fmt.Sprintf("select column_name, column_comment, data_type, column_key, extra from information_schema.columns where table_schema='%s' and table_name= '%s'", config.Database, tableName))
	queryColumns, _ = rows.Columns()
	values := make([]sql.RawBytes, len(queryColumns))
	scans := make([]interface{}, len(queryColumns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]string)
		for i, col := range values {
			each[queryColumns[i]] = string(col)
		}
		columns = append(columns, each)
	}
	return
}

// GetTables 查询表
func GetTables() (tables []map[string]string) {
	var (
		colSql strings.Builder
	)

	colSql.WriteString(fmt.Sprintf("select table_name, table_comment from information_schema.tables where table_schema='%s'", config.Database))
	if !config.IsAllTables && config.Table != "" {
		colSql.WriteString(fmt.Sprintf(" and table_name ='%s'", config.Table))
	}

	rows, _ := Db.Query(colSql.String())
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	// 所有表
	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]string)
		for i, col := range values {
			each[columns[i]] = string(col)
		}
		tables = append(tables, each)
	}
	return
}

func Close() {
	Db.Close()
}
