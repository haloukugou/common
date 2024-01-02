package model

import "time"

type Apk struct {
	Id        int       `json:"id"`
	File      string    `json:"file"`
	Version   string    `json:"version"`
	IsForce   int       `json:"isForce"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (table *Apk) TableName() string {
	return "apk"
}
