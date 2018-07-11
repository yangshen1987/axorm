package Mysql

import (
	"database/sql"
	"updateServer/Conf"
	"updateServer/Db"
)




type MysqlOBj struct {
	Db.Base
	sql string
	Tx *sql.Tx
	Stmt *sql.Stmt
	result struct{}
} 
func init()  {
	Mysql := new(MysqlOBj)
	Mysql.Pool = Conf.MYSQLDSN
	db, err := sql.Open("mysql", Conf.MYSQLDSN)
	if err!= nil{
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err.Error())
	}
}
func (m *MysqlOBj)CreateConnection()  {
	mysql, err := sql.Open("mysql", Conf.MYSQLDSN)
	if err != nil{
		panic(err.Error())
	}
	defer mysql.Close()
	err = mysql.Ping()
	if err!= nil{
		panic(err.Error())
	}
	m.connection = mysql
}

func (m *MysqlOBj)Prepare(sql string)  {
	m.sql = sql
}

func (m *MysqlOBj)Execut(param []string)*sql.Rows  {
	m.FormatSqlByParams(param)
	result, err := m.connection.Query(m.sql)
	if err != nil{
		panic(err.Error())
	}
	return result
}

func (m *MysqlOBj)Receive()  {
	//
}

func (m *MysqlOBj)Fetch() struct{}  {
	return m.result
}

func (m *MysqlOBj)GetLastInsertId() int  {
	return 1
}
func (m *MysqlOBj)GetAffectedRows() int  {
	return 2
}
func (m *MysqlOBj)BeginTransaction()  {
	tx,err := m.connection.Begin()
	if err != nil{
		panic(err.Error())
	}
	m.Tx = tx
}
func (m *MysqlOBj)Rollback()  {
	m.Tx.Rollback()

}
func (m *MysqlOBj)Commit()  {
	m.Tx.Commit()

}
//暂时没有 找到好办法实现他
func (m *MysqlOBj)SelectDb(Db string)  {

}
func (m *MysqlOBj)SetDefer(defered bool)  {

}
func (m *MysqlOBj)Reconnect()  {
	m.CreateConnection()
}
func (m *MysqlOBj)Check() bool {
	err := m.connection.Ping()
	if err != nil{
		return false
	}
	return true
}
func (m *MysqlOBj)Destroy()  {
	m.sql = ""

}
func (m *MysqlOBj)FormatSqlByParams(param []string)  {
	if m.Tx == nil{
		stmt, err := m.connection.Prepare(m.sql)
		if err != nil{
			panic(err.Error())
		}
		m.Stmt = stmt

	}else {
		stmt, err := m.Tx.Prepare(m.sql)
		if err != nil{
			panic(err.Error())
		}
		m.Stmt = stmt

	}
}
func (m *MysqlOBj)TransferQuestionMark() {

}
func (m *MysqlOBj)GetSql() string  {
	return m.sql
}

