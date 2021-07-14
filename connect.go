package main

import "database/sql"


func connect()  *sql.DB{

	var err error
	var db *sql.DB

	db,err = sql.Open("postgres", dataSourceName)

	checkErr(err)

	return db
}