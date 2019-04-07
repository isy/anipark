package model

import (
	"github.com/jinzhu/gorm"
)

type Tweet struct {
	gorm.Model
	Text string `json:"text" gorm:"not null" sql:"type:text"`
}
