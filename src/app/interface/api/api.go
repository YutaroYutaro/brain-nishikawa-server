package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"src/app/interface/api/handler"
	"src/app/registry"
	"src/app/usecase"
)

func Api(ctn *registry.Container, r *gin.Engine, db *gorm.DB) {
	meetingHandler := handler.NewMeetingHandler(ctn.Resolve("meeting-usecase").(usecase.MeetingUseCase))

	r.GET("/meeting", wrapper(meetingHandler.ListMeeting, db))
	r.GET("/meeting/:id", wrapper(meetingHandler.GetMeetingByID, db))

}

func wrapper(fn func(*gin.Context, *gorm.DB), db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fn(ctx, db)
	}
}
