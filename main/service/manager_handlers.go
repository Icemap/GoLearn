package service

import (
	"net/http"
	"strconv"
	"GOLEARN/main/dao"
	"encoding/json"
)

func InsertManager(responseWriter http.ResponseWriter, request *http.Request)  {
	managerId := dao.InsertManager(
		request.FormValue("name"),
		request.FormValue("account"),
		request.FormValue("password"))
	responseWriter.Write([]byte (strconv.Itoa(managerId)))
}

func DeleteManager(responseWriter http.ResponseWriter, request *http.Request)  {
	id ,_ :=strconv.Atoi(request.FormValue("id"))
	dao.DeleteManager(id)
	responseWriter.Write([]byte ("true"))
}

func SelectManager(responseWriter http.ResponseWriter, request *http.Request)  {
	data := dao.SelectAllManager()
	strJson, _ := json.Marshal(data)
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write(strJson)
}