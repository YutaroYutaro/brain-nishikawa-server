package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (h *meetingHandler) ListMeeting(c *gin.Context) {
	meetings, err := h.meetingUseCase.ListMeeting()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(
		200,
		meetings,
	)
}

func (h *meetingHandler) GetMeetingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	meeting, err := h.meetingUseCase.GetMeetingByID(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(
		200,
		meeting,
	)
}
