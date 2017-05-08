package model

type User struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	Age         string `xorm:"VARCHAR(11)"`
	Income      string `xorm:"VARCHAR(11)"`
	Homeaddress string `xorm:"VARCHAR(100)"`
	Workaddress string `xorm:"VARCHAR(100)"`
	Projid      int    `xorm:"INT(11)"`
	Vistamatrix string `xorm:"VARCHAR(1000)"`
	Vistaresult string `xorm:"VARCHAR(200)"`
	Hold        string `xorm:"VARCHAR(200)"`
}
