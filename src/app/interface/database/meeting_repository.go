package database

import (
	"github.com/jinzhu/gorm"
	"src/app/domain/model"
)

type meetingRepository struct {
	db *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) *meetingRepository {
	return &meetingRepository{
		db: db,
	}
}

func (r *meetingRepository) FindAll() ([]*model.Meeting, error) {
	var meetings []Meeting
	r.db.Table("meeting").Select("id, title").Find(&meetings)

	modelMeetings := make([]*model.Meeting, len(meetings))
	i := 0
	for _, meeting := range meetings {
		modelMeetings[i] = model.NewMeeting(meeting.ID, meeting.Title)
		i++
	}

	return modelMeetings, nil
}

func (r *meetingRepository) FindByID(id int) (*model.Meeting, error) {
	var meeting Meeting
	r.db.Table("meeting").Select("id, title").Where("id = ?", id).Find(&meeting)

	modelMeeting := model.NewMeeting(meeting.ID, meeting.Title)

	return modelMeeting, nil
}

type Meeting struct {
	ID    int
	Title string
}
