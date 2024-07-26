// Package dbmodel Code generated by sql2struct. https://github.com/dxc0522/sql2struct
package dbmodel

import "time"

type ChargingData struct {
	ID         int64     `gorm:"column:id" json:"id,omitempty"`
	Waterpower float64   `gorm:"column:waterpower" json:"waterpower,omitempty"`
	Allpower   float64   `gorm:"column:allpower" json:"allpower,omitempty"`
	Allmoney   float64   `gorm:"column:allmoney" json:"allmoney,omitempty"`
	Initpower  float64   `gorm:"column:initpower" json:"initpower,omitempty"`
	Lastpower  float64   `gorm:"column:lastpower" json:"lastpower,omitempty"`
	Difference float64   `gorm:"column:difference" json:"difference,omitempty"`
	Price      int64     `gorm:"column:price" json:"price,omitempty"`
	Date       time.Time `gorm:"column:date" json:"date,omitempty"`
	UUID       string    `gorm:"column:uuid" json:"uuid,omitempty"`
}

// TableName the name of table in database
func (t *ChargingData) TableName() string {
	return "charging_data"
}
