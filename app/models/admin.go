package models

import "time"

type Admin struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Account   string    `json:"account"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 设置表名
func (table *Admin) TableName() string {
	return "admin"
}
