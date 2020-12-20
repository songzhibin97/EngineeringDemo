package data

import (
	"Songzhibin/EngineeringDemo/internal/biz"

	"gorm.io/gorm"
)

type LogicData struct {
	db *gorm.DB
}

func NewLogicData(db *gorm.DB) *LogicData {
	return &LogicData{db: db}
}

func (l *LogicData) Login(username, password string) bool {
	u := &biz.User{}
	if l.db.Model(u).First("username = ? AND password = ?", username, password) != nil {
		return false
	}
	return true

}

func (l *LogicData) Registered(user *biz.User) bool {
	u := &biz.User{}
	if l.db.Model(u).First("username = ?", user.Username).Error != nil {
		return false
	}
	if l.db.Create(user).Error != nil {
		return false
	}
	return true
}
