package main

type GormUser struct {
	ID       uint `gorm:"primarykey"`
	Username string
	Password string
}

type Tabler interface {
	TableName() string
}

func (GormUser) TableName() string {
	return "dbo.users"
}
