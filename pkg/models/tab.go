package models

import (
	"time"

	"github.com/Khotso-Bore/tab/pkg/config"
)

type Tab struct {
	Id         int64
	CreditorId int64
	DebtorId   int64
	Reason string
	Amount     float64
	CreatedAt time.Time
}

func GetTabsByCreditorDebtorIds(creditorId int64, debtorId int64) []Tab{

	db := config.GetDB()
	
	query := "SELECT id,reason,createdAt,amount FROM tabs WHERE creditorId = ? AND debtorId = ? "
	rows,err := db.Query(query,creditorId,debtorId)
	if err != nil{
		panic(err)
	}

	defer rows.Close()
	
	
	var tabs []Tab
	for rows.Next() {

		var t Tab
		err := rows.Scan(&t.Id,&t.Reason,&t.CreatedAt,&t.Amount)
		if err != nil{
			panic(err)
		}

		
		tabs = append(tabs, t)

	}

	return tabs

}


func CreateTab(tab Tab) Tab{
	db := config.GetDB()

	tab.CreatedAt = time.Now()
	query := "INSERT INTO tabs (creditorId,debtorId,reason,createdAt,amount) VALUES (?,?,?,?,?)"
	res,err := db.Exec(query,tab.CreditorId,tab.DebtorId,tab.Reason,tab.CreatedAt,tab.Amount)
	if err != nil{
		panic(err)
	}

	tab.Id,_ = res.LastInsertId();
	return tab
}

func DeleteTab(id int64) {
	db := config.GetDB()

	query := "DELETE FROM users WHERE id = ?"
	_,err := db.Exec(query,id)
	if err != nil{
		panic(err)
	}
}

