package Db

import (
	"database/sql"
	"updateServer/Conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect()  {
	db, err := sql.Open("mysql", Conf.MYSQLDSN)
	if err!= nil{
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err.Error())
	}
	rows, err := db.Query("select * from ims_store_application_branch")
	if err != nil{
		panic(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil{
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values{
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}