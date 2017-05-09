package dao

import "main/model"

func SelectAllManager() []model.Manager {
	data := make([]model.Manager, 0)
	DBEngine.Find(&data)
	return data
}

func DeleteManager(Id int){
	var manager model.Manager
	DBEngine.Id(Id).Delete(&manager)
}

func InsertManager(Name string, Username string, Password string) int {
	var manager model.Manager
	manager.Name = Name
	manager.Username = Username
	manager.Password = Password
	DBEngine.Insert(&manager)
	return manager.Id
}