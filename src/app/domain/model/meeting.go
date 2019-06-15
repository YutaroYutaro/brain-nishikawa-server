package model

type Meeting struct {
	ID    int
	Title string
}

func NewMeeting(id int, title string) *Meeting {
	return &Meeting{
		ID:    id,
		Title: title,
	}
}

func (m *Meeting) GetID() int {
	return m.ID
}

func (m *Meeting) GetTitle() string {
	return m.Title
}
