package model

type Manager struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"VARCHAR(20)"`
	Username string `xorm:"VARCHAR(50)"`
	Password string `xorm:"VARCHAR(50)"`
}
