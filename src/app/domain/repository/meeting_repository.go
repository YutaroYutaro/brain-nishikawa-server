package repository

import (
	"github.com/jinzhu/gorm"
	"src/app/domain/model"
)

type MeetingRepository interface {
	FindAll(*gorm.DB) ([]*model.Meeting, error)
	FindByID(*gorm.DB, int) (*model.Meeting, error)
	//Save(*model.Meeting) error
}
