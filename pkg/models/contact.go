package models

import (
	"fmt"

	"github.com/Khotso-Bore/tab/pkg/config"
)

type Contact struct {
	UserId    int64
	ContactId int64
}

func GetUserContacts(userId int64) []int64 {

	db := config.GetDB()

	fmt.Println(userId)
	query := "SELECT contactId FROM contacts WHERE userId = ?"
	rows,err := db.Query(query,userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	


	var contactIds []int64
	for rows.Next(){

		var id int64
		rows.Scan(&id)
		contactIds = append(contactIds, id)
	}

	return contactIds
}

func CreateContact(userid int64,contactId int64) {

	db := config.GetDB();

	query := "INSERT INTO contacts (userId,conatctId) VALUES (?,?)"
	_,err := db.Exec(query,userid,contactId)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec(query,contactId,userid)
	if err != nil {
		panic(err)
	}
}  