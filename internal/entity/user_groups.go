package entity

type UserGroups struct {
	Id      int `gorm:"primaryKey"`
	GroupId int
	UserId  int
}
