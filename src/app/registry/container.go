package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di"
	"src/app/domain/service"
	"src/app/interface/database"
	"src/app/usecase"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "meeting-usecase",
			Build: buildMeetingUseCase,
			Close: func(obj interface{}) error {
				return obj.(*gorm.DB).Close()
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil

}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildMeetingUseCase(ctn di.Container) (interface{}, error) {
	db, err := gorm.Open("mysql", "root:password@tcp(db-server:3306)/test_db")
	if err != nil {
		return nil, err
	}
	repo := database.NewMeetingRepository(db)
	service := service.NewMeetingService(repo)

	return usecase.NewMeetingUseCase(repo, service), nil
}
