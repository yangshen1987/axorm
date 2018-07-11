package Db

type Result interface {
	getResult(param ...interface{})
} 
