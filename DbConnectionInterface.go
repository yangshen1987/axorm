package Db

import "database/sql"

type Mysql interface {
	CreateConnection()
	Prepare(sql string)
	Execut(param []string) *sql.Rows
	Receive()
	Fetch() struct{}
	GetLastInsertId() int
	GetAffectedRows() int
	BeginTransaction()
	Rollback()
	Commit()
	SelectDb(Db string)
	SetDefer(defered bool)
	Reconnect()
	Check() bool
	Destroy()
	FormatSqlByParams(param []string)
	TransferQuestionMark()
	GetSql() string
}
type Base struct {
	connection *sql.DB
	Pool string
}
