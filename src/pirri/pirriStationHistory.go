package main

import (
	//	"fmt"
	"time"
)

type StationHistory struct {
	ID         int       `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	StationID  int       `gorm:"not null"`
	ScheduleID int       `sql:"DEFAULT:0" gorm:"not null"`
	Duration   int       `sql:"DEFAULT:0" gorm:"not null"`
	StartTime  time.Time `sql:"DEFAULT:current_timestamp"`
}