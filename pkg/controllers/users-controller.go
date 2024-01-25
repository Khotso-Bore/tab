package controllers

import (
	"net/http"
	"strconv"

	"github.com/Khotso-Bore/tab/pkg/models"
	"github.com/Khotso-Bore/tab/pkg/utils"
	"github.com/gorilla/mux"
)


func GetUser(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	id,_ := strconv.ParseInt(vars["id"],0,0)

	user,err := models.GetUser(id)

	if err != nil{
		
		utils.Response(&w,http.StatusNotFound,nil)
		return
	}

	utils.Response(&w,http.StatusOK,user)
}

func GetUsers(w http.ResponseWriter,r *http.Request){
	
	users := models.GetUsers()
	utils.Response(&w,http.StatusOK,users)
	
}

func CreateUser(w http.ResponseWriter,r *http.Request){

	var user models.User
	utils.ParseBody(r,&user)

	newUser := models.CreateUser(user)

	utils.Response(&w,http.StatusOK,newUser)
	
}

func UpdateUser(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	id,_ := strconv.ParseInt(vars["id"],0,0)

	user,err := models.GetUser(id)
	if err != nil{
		
		utils.Response(&w,http.StatusNotFound,nil)
		return
	}

	var updatedUser models.User
	utils.ParseBody(r,&updatedUser)

	if updatedUser.Cellnumber != ""{
		user.Cellnumber = updatedUser.Cellnumber
	}

	if updatedUser.Username != ""{
		user.Username = updatedUser.Username
	}

	if updatedUser.Paysharpid != ""{
		user.Paysharpid = updatedUser.Paysharpid
	}

	user = models.UpdateUser(user)
	utils.Response(&w,http.StatusOK,user)

}

func DeleteUser(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	id,_ := strconv.ParseInt(vars["id"],0,0)

	models.DeleteUser(id)

	utils.Response(&w,http.StatusOK,"deleted")

}

func AddContact(w http.ResponseWriter,r *http.Request){
	
	var contact models.Contact
	utils.ParseBody(r,&contact)

	models.CreateContact(contact.UserId,contact.ContactId)

	utils.Response(&w,http.StatusOK,"ok")

}

func GetContacts(w http.ResponseWriter,r *http.Request){

	vars := mux.Vars(r)
	id,_ := strconv.ParseInt(vars["id"],0,0)

	contactIds := models.GetUserContacts(id)

	users := models.GetUsersInList(contactIds)

	utils.Response(&w,http.StatusOK,users)	
}

func DepositFunds(w http.ResponseWriter,r *http.Request){

	var newBalance struct {
		UserId int64
		Amount float64
	}

	newBalance.UserId = 9000
	utils.ParseBody(r,&newBalance)

	user,err := models.GetUser(newBalance.UserId)
	if err != nil{
		
		utils.Response(&w,http.StatusNotFound,nil)
		return
	}

	user.Balance = newBalance.Amount + user.Balance
	user = models.UpdateUser(user)
	utils.Response(&w,http.StatusOK,user)

}

