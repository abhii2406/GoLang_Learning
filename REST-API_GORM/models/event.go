package models

import (
	"errors"
	"learning/rest-api_gorm/db"
	"time"
)

type Event struct {
	Id          int64     `gorm:"primaryKey"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
	User        User `gorm:"foreignKey:UserId" json:"-" binding:"-"`
}

// var events = []Event{}

func (e *Event) Save() error {
	return db.DB.Create(&e).Error
}

func GetAllevents() ([]Event, error) {
	var event []Event
	err := db.DB.Preload("User").Find(&event).Error
	return event, err
}

func GetEventById(id int64) (*Event, error) {

	var event Event
	err := db.DB.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (event Event) Update() error {

	res := db.DB.Model(&Event{}).Where("id = ?", event.Id).Updates(Event{Name: event.Name, Description: event.Description, Location: event.Location, DateTime: event.DateTime})
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return res.Error
}

func (event Event) Delete() error {
	res := db.DB.Delete(&Event{}, event.Id)
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return res.Error

}
