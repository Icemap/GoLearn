package model

type Proj struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Title   string `xorm:"VARCHAR(50)"`
	Content string `xorm:"VARCHAR(200)"`
	Belong  int    `xorm:"INT(11)"`
}
