package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"src/app/usecase"
	"strconv"
)

type meetingHandler struct {
	meetingUseCase usecase.MeetingUseCase
}

func NewMeetingHandler(meetingUseCase usecase.MeetingUseCase) *meetingHandler {
	return &meetingHandler{
		meetingUseCase: meetingUseCase,
	}
}

func (h *meetingHandler) ListMeeting(c *gin.Context, db *gorm.DB) {
	meetings, err := h.meetingUseCase.ListMeeting(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(
		200,
		meetings,
	)
}

func (h *meetingHandler) GetMeetingByID(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	meeting, err := h.meetingUseCase.GetMeetingByID(db, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(
		200,
		meeting,
	)
}
