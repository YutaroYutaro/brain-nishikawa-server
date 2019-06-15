package service

import "src/app/domain/repository"

type MeetingService struct {
	repo repository.MeetingRepository
}

func NewMeetingService(repo repository.MeetingRepository) *MeetingService {
	return &MeetingService{
		repo: repo,
	}
}
