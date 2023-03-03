package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "localhost"
	port     = 1433
	user     = "testgo"
	password = "1514Vorsha@"
	database = "mydb"
)

func CreateEmployee(db *sql.DB, username string, password string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO dbo.users (username, password) VALUES ('%s','%s');",
		username, password)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

// ReadEmployees read all employees
func ReadEmployees(db *sql.DB) (int, error) {
	tsql := fmt.Sprintf("SELECT id, username, password FROM dbo.users;")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("id: %d, username: %s, password: %s\n", user.Id, user.Username, user.Password)
		count++
	}
	return count, nil
}

// UpdateEmployee update an employee's information
func UpdateEmployee(db *sql.DB, username string, password string) (int64, error) {
	tsql := fmt.Sprintf("UPDATE dbo.users SET password = '%s' WHERE username= '%s'",
		password, username)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

// DeleteEmployee delete an employee from database
func DeleteEmployee(db *sql.DB, username string) (int64, error) {
	tsql := fmt.Sprintf("DELETE FROM dbo.users WHERE Name='%s';", username)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}
