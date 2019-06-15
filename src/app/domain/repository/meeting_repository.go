package repository

import (
	"src/app/domain/model"
)

type MeetingRepository interface {
	FindAll() ([]*model.Meeting, error)
	FindByID(int) (*model.Meeting, error)
	//Save(*model.Meeting) error
}
