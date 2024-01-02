package model

import (
	"time"
)

type User struct {
	Id        uint64    `gorm:"column:id;auto_increment;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(50)" json:"name"`
	Account   string    `gorm:"column:account;type:varchar(50)" json:"account"`
	Title     string    `gorm:"column:title;type:varchar(50)" json:"title"`
	Salt      string    `gorm:"column:salt;type:varchar(50)" json:"salt,omitempty"`
	Password  string    `gorm:"column:password;type:varchar(50)" json:"password,omitempty"`
	Mail      string    `gorm:"column:mail;type:varchar(60)" json:"mail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Source    int       `json:"source"`
}

// TableName 设置表名
func (table *User) TableName() string {
	return "user"
}
