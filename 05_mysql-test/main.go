package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//Standard opening
	db, err := sql.Open("mysql", "root:XLInts1991#@tcp(127.0.0.1:3306)/go_learning")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//Ping it to check for errors first
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("No errors pinging database. Proceed..")
	}
	//Begin statements
	var (
		id    int
		name  string
		email string
	)
	rows, err := db.Query("select * from go_users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &id, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, email)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO go_users(user_id, username, user_email) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(2, "OctogonKnife", "stabbing.westworld@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	fmt.Println("Database Closed")
}
