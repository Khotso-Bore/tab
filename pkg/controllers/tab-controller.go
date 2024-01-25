package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Khotso-Bore/tab/pkg/models"
	"github.com/Khotso-Bore/tab/pkg/utils"
	"github.com/gorilla/mux"
)

func GetTabForUserContact(w http.ResponseWriter, r* http.Request){

	vars := mux.Vars(r)
	userId,_ := strconv.ParseInt(vars["userid"],0,0)
	contactId,_ := strconv.ParseInt(vars["contactid"],0,0)

	_,err := models.GetUser(userId)
	if err != nil {
		utils.Response(&w,http.StatusOK,fmt.Sprintf("user with id %v does not exist",userId))
		return;
	}

	_,err = models.GetUser(contactId)
	if err != nil {
		utils.Response(&w,http.StatusOK,fmt.Sprintf("user with id %v does not exist",contactId))
		return;
	}

	userTabs := models.GetTabsByCreditorDebtorIds(userId,contactId)
	contactTabs := models.GetTabsByCreditorDebtorIds(contactId,userId)

	userTabs = append(userTabs, contactTabs...)

	utils.Response(&w,http.StatusOK,userTabs)

}

func CreateTab(w http.ResponseWriter, r* http.Request){

	var tab models.Tab
	utils.ParseBody(r,&tab)

	_,err := models.GetUser(tab.CreditorId)
	if err != nil {
		utils.Response(&w,http.StatusNotFound,fmt.Sprintf("user with id %v does not exist",tab.CreditorId))
		return;
	}

	_,err = models.GetUser(tab.DebtorId)
	if err != nil {
		utils.Response(&w,http.StatusNotFound,fmt.Sprintf("user with id %v does not exist",tab.DebtorId))
		return;
	}

	if tab.Amount <= 0 {
		utils.Response(&w,http.StatusBadRequest,"Amount must be greater than 0")
		return;
	}

	if tab.Reason == "" {
		utils.Response(&w,http.StatusBadRequest,"Please provide a reason for opening this tab")
		return;
	}

	tab = models.CreateTab(tab)
	utils.Response(&w,http.StatusOK,tab)
}

func PayTab(w http.ResponseWriter,r *http.Request){

	var tab models.Tab
	utils.ParseBody(r,&tab)

	contact,err := models.GetUser(tab.DebtorId)
	if err != nil {
		utils.Response(&w,http.StatusNotFound,fmt.Sprintf("user with id %v does not exist",tab.DebtorId))
		return;
	}

	user,err := models.GetUser(tab.CreditorId)
	if err != nil {
		utils.Response(&w,http.StatusNotFound,fmt.Sprintf("user with id %v does not exist",tab.CreditorId))
		return;
	}

	if user.Balance < tab.Amount {
		utils.Response(&w,http.StatusNotFound,"user lacks the funds to make this payment")
		return;
	}

	if tab.Amount <= 0 {
		utils.Response(&w,http.StatusBadRequest,"Amount must be greater than 0")
		return;
	}

	tab.Reason = "PAYMENT"

	tab = models.CreateTab(tab)

	user.Balance = user.Balance - tab.Amount
	models.UpdateUser(user)

	contact.Balance = tab.Amount + contact.Balance
	models.UpdateUser(contact)

	utils.Response(&w,http.StatusOK,"payment succesful")

}