package model

type Vista struct {
	Id      int     `xorm:"not null pk autoincr INT(11)"`
	Url     string  `xorm:"VARCHAR(500)"`
	Content string  `xorm:"VARCHAR(200)"`
	Belong  int     `xorm:"INT(11)"`
	Lon     float64 `xorm:"DOUBLE"`
	Lat     float64 `xorm:"DOUBLE"`
}
