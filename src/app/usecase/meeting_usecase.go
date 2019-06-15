package usecase

import (
	"src/app/domain/model"
	"src/app/domain/repository"
	"src/app/domain/service"
)

type MeetingUseCase interface {
	ListMeeting() ([]*Meeting, error)
	GetMeetingByID(int) (*Meeting, error)
	//RegisterMeeting(title string) error
}

type meetingUseCase struct {
	repo    repository.MeetingRepository
	service *service.MeetingService
}

func NewMeetingUseCase(repo repository.MeetingRepository, service *service.MeetingService) *meetingUseCase {
	return &meetingUseCase{
		repo:    repo,
		service: service,
	}
}

func (m *meetingUseCase) ListMeeting() ([]*Meeting, error) {
	meetings, err := m.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return toMeeting(meetings), nil
}

func (m *meetingUseCase) GetMeetingByID(id int) (*Meeting, error) {
	meeting, err := m.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &Meeting{
		ID:    meeting.ID,
		Title: meeting.Title,
	}, nil
}

//func (m *meetingUseCase) RegisterMeeting(title string) error {
//	meeting := model.NewMeeting(0, title)
//	if err := m.repo.Save(meeting); err != nil {
//		return err
//	}
//
//	return nil
//}

type Meeting struct {
	ID    int
	Title string
}

func toMeeting(meetings []*model.Meeting) []*Meeting {
	res := make([]*Meeting, len(meetings))
	for i, meeting := range meetings {
		res[i] = &Meeting{
			ID:    meeting.ID,
			Title: meeting.Title,
		}
	}

	return res
}
