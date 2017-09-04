package dao

import "GOLEARN/main/model"

func InsertVista(Url string, Content string, Belong int, Lon float64, Lat float64) int {
	var vista model.Vista
	vista.Url = Url
	vista.Content = Content
	vista.Belong = Belong
	vista.Lon = Lon
	vista.Lat = Lat
	DBEngine.Insert(&vista)

	return vista.Id
}
