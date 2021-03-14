package models

import (
	"time"
)

type Calculations struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Num1 string `gorm:"size:255;not null" json:"num1"`
	Operator string `gorm:"size:25;not null" json:"operator"`
	Num2 string `gorm:"size:255;not null" json:"num2"`
	Result string `gorm:"size:255;not null" json:"result"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (c *Calculations) TableName() string {
	return "calculations"
}
