package service

import (
	"fmt"
	"time"

	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/storage"
)

type EventService struct {
	storage storage.Storage
}

// func NewEventService(storage storage.Storage) *EventService {
// 	return &EventService{
// 		storage: storage, 
// 	}
// }

func (s *EventService) GetEventForDay(day string) ([]*events.Event, error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, day); err != nil {
		return nil, err
	} else {
		return s.storage.GetEventForDay(t)
	}
	
}

func (s *EventService) GetEventForWeek(dayStart string) ([]*events.Event, error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, dayStart); err != nil {
		return nil, err
	} else {
		return s.storage.GetEventForWeek(t)
	}
}

func (s *EventService) GetEventForMonth(dayStart string) ([]*events.Event, error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, dayStart); err != nil {
		return nil, err
	} else {
		return s.storage.GetEventForMonth(t)
	}
}

func (s *EventService) Add(title string, date string) (*events.Event, error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, date); err != nil {
		return nil, err
	} else {
		return s.storage.Add(title, t)
	}
	
}

func (s *EventService) Delete(day string, title string) (*events.Event,error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, day); err == nil {
		return nil, err
	} else {
		return s.storage.Delete(t, title)
	}
	
}

func (s *EventService) Update(day string, title, newTitle string) (*events.Event,error) {
	layout := "2006-01-02"
	if t, err := time.Parse(layout, day); err != nil {
		return nil, err
	} else {
		if newTitle == "" {
			return nil, fmt.Errorf("Пустая заметка %s", title)
		}
		return s.storage.Update(t, title, newTitle)
	}
}