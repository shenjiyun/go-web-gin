package model

type Users struct {
	ID uint  `gorm:primaryKey`
	Nickname string
}

func (Users) TableName() string {
	return "users"
}
