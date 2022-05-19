package service

import (
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/storage"
)

type ServiceInterface interface {
	GetEventForDay(day string) ([]*events.Event, error)
	GetEventForWeek(dayStart string) ([]*events.Event, error)
	GetEventForMonth(dayStart string) ([]*events.Event, error)
	Add(title string, date string) (*events.Event,error)
	Delete(day string, title string) (*events.Event,error)
	Update(day string, title, newTitle string) (*events.Event,error)
}

type Service struct {
	ServiceInterface
}

func NewService(storage *storage.Storage) *Service {
	return &Service{&EventService{storage}}
}