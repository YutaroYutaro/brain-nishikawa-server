package api

import (
	"github.com/gin-gonic/gin"
	"src/app/interface/api/handler"
	"src/app/registry"
	"src/app/usecase"
)

func Api(ctn *registry.Container, r *gin.Engine) {
	meetingHandler := handler.NewMeetingHandler(ctn.Resolve("meeting-usecase").(usecase.MeetingUseCase))

	r.GET("/meeting", meetingHandler.ListMeeting)
	r.GET("/meeting/:id", meetingHandler.GetMeetingByID)

}
