package util

import "database/sql"

// 执行sql
func Execute(db *sql.DB, sql string, args ...interface{}) (sql.Result, error) {
	tx, err := db.Begin()
	if err != nil {
		db.Close()
		return nil, err
	}
	result, err := tx.Exec(sql, args...)
	if err != nil {
		db.Close()
		return nil, err
	}
	tx.Commit()
	db.Close()
	return result, err
}
