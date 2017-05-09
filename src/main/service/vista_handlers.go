package service

import (
	"io"
	"strconv"
	"main/dao"
	"net/http"
	"time"
	"os"
	"main/utils"
)

func InsertVista(responseWriter http.ResponseWriter, request *http.Request)  {
	var filename string

	file, _, err := request.FormFile("icon")
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		return
	}

	defer file.Close()
	filename = strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	f,err := os.Create(utils.ConfUtils().VistaImageSavePath + filename)
	defer f.Close()
	io.Copy(f,file)

	belong, _ := strconv.Atoi(request.FormValue("belong"))
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)

	vistaId := dao.InsertVista(
		filename,
		request.FormValue("content"),
		belong,
		lon,
		lat)
	responseWriter.Write([]byte (strconv.Itoa(vistaId)))
}