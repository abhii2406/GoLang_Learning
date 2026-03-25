// package db

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDB() {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "api.db")

// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)

// 	}

// 	DB.SetMaxOpenConns(10) // Set the maximum number of open connections to the database
// 	DB.SetMaxIdleConns(5)  // Set the maximum number of idle connections in the pool  (means the number of connections that can be kept open and reused)

// 	createTables()
// }

// func createTables() {
// 	createUsersTable := `
// 		create table if not exists users(
// 		id integer primary key autoincrement,
// 		email text not null unique,
// 		password text not null
// 		)
// 	`
// 	_, err := DB.Exec(createUsersTable)
// 	if err != nil {
// 		panic("could not create users table")
// 	}
// 	createEventsTable := `
// CREATE TABLE IF NOT EXISTS events (
// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	name TEXT NOT NULL,
// 	description TEXT NOT NULL,
// 	location TEXT NOT NULL,
// 	date_time DATETIME NOT NULL,
// 	user_id INTEGER,
// 	foreign key(user_id) references user(id)
// );
// `
// 	_, err = DB.Exec(createEventsTable)
// 	if err != nil {
// 		panic("could not create event table")
// 	}
// }

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/practice_go?parseTime=true")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create event table")
	}
}
