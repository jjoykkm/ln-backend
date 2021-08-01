package utility

import (
	"database/sql"
	"fmt"
	"log"
)


func GetCountTable(db *sql.DB,status, tableName, field, condition string) int {
	var count int

	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE status_id = $1 AND %s ", field, tableName, condition)
	err := db.QueryRow(sql, status).Scan(
		&count ,
	)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func SelectData(db *sql.DB, selectList, tableName, condition, joinTable, joinKey, orderBy, status string, offset int, limit int) *sql.Rows {
	if selectList == "" && tableName == "" {
		return nil
	}

	sql :=  fmt.Sprintf("SELECT %s FROM %s", selectList, tableName)
	if joinTable != "" && joinKey != "" {
		sqlJoin := fmt.Sprintf(" INNER JOIN %s ON %s WHERE %s.status_id = '%s' ", joinTable, joinKey, tableName, status)
		sql = sql + sqlJoin
	}
	if joinTable == "" && joinKey == "" {
		sqlStatus := fmt.Sprintf(" WHERE status_id = '%s' ", status)
		sql = sql + sqlStatus
	}

	if condition != "" {
		sqlOrder := fmt.Sprintf(" AND %s", condition)
		sql = sql + sqlOrder
	}
	if orderBy != "" {
		sqlOrder := fmt.Sprintf(" ORDER BY %s", orderBy)
		sql = sql + sqlOrder
	}
	if offset > 0 {
		sqlOffset := fmt.Sprintf(" OFFSET %d", offset)
		sql = sql + sqlOffset
	}
	if limit > 0 {
		sqlLimit := fmt.Sprintf(" LIMIT %d", limit)
		sql = sql + sqlLimit
	}
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}

func SelectDataManual(db *sql.DB, sql string) *sql.Rows {
	if sql == "" {
		return nil
	}
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}