package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Khotso-Bore/tab/pkg/config"
)

var db *sql.DB

var query string = `
CREATE TABLE users (
	id INT AUTO_INCREMENT,
	username TEXT NOT NULL,
	cellnumber TEXT NOT NULL,
	paysharpid TEXT NOT NULL,
	balance DOUBLE,
	PRIMARY KEY (id)
);
`

type User struct {
	Id int64
	Username string
	Cellnumber string
	Paysharpid string
	Balance float64
}



func GetUser(id int64) (User,error){

	db := config.GetDB()
	var user User;
	
	query := `SELECT * FROM users WHERE id = ? `
	err := db.QueryRow(query,id).Scan(&user.Id,&user.Username,&user.Cellnumber,&user.Paysharpid,&user.Balance)
	if err != nil {
		return user,err

	}
	
	return user,nil

}

func GetUsers() []User{

	db := config.GetDB()

	query := `SELECT * FROM users`
	rows,err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close();

	var users []User
	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id,&u.Username,&u.Cellnumber,&u.Paysharpid,&u.Balance)
		if err != nil {
			panic(err)
		}
		users =  append(users,u)
	}
	
	return users

}

func CreateUser(user User) User{

	db := config.GetDB()
	query := `INSERT INTO users (username,cellnumber,paysharpid,balance) VALUES (?,?,?,?)`
	res,err := db.Exec(query,&user.Username,&user.Cellnumber,&user.Paysharpid,&user.Balance)
	if err != nil {
		panic(err)
	}

	user.Id,err = res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return user

}

func UpdateUser(user User) User{
	
	db := config.GetDB()
	query := `
	Update users SET 
	username = ?,
	cellnumber = ?,
	paysharpid = ?,
	balance = ?
	WHERE id = ?
	`

	_,err := db.Exec(query,user.Username,user.Cellnumber,user.Paysharpid,user.Balance,user.Id)
	if err != nil {
		panic(err)
	}

	return user;
}

func DeleteUser(id int64) {
	
	db := config.GetDB()
	query := `DELETE FROM users WHERE id = ?`
	_,err := db.Exec(query,id)
	if err != nil {
		panic(err)
	}
}

func GetUsersInList(ids []int64) []User {

	db := config.GetDB()

	fmt.Println("query")

	placeholders := make([]string,len(ids))
	for i,_ := range ids {
		placeholders[i] = "?"
	}
	args := make([]interface{},len(ids))
	for i,id := range ids {
		args[i] = id
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE id in (%v)",strings.Join(placeholders,","))
	rows,err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close();
	fmt.Println("rows")

	var users []User
	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id,&u.Username,&u.Cellnumber,&u.Paysharpid,&u.Balance)
		if err != nil {
			panic(err)
		}
		users = append(users,u)
	}
	
	return users

}



