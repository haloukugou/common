package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Follow struct {
	Id             uint64     `gorm:"column:id;auto_increment;primary_key;type:bigint" json:"id"`
	FollowPerson   uint64     `gorm:"column:follow_person;type:bigint" json:"follow_person"`
	FollowedPerson uint64     `gorm:"column:followed_person;type:bigint" json:"followed_person"`
	CreatedAt      *LocalTime `json:"created_at"`
	UpdatedAt      *LocalTime `json:"updated_at"`
}

// TableName 设置表名
func (table *Follow) TableName() string {
	return "follow"
}

type LocalTime time.Time

// MarshalJSON 在转json时转成定义得时间格式
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value 值插入时间戳到mysql需要这个函数
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
