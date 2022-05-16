package storage

import (
	"fmt"
	"time"

	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
)


func (s *Storage) GetEventForDay(day time.Time) ([]*events.Event, error) {
	s.Lock()
	for date, event := range s.events {
		if day == date {
			return event, nil			
		} 
	}
	defer s.Unlock()
	return nil, fmt.Errorf("Нет заметок для этого дня")
}

func (s *Storage) GetEventForWeek(dayStart time.Time) ([]*events.Event, error) {
	s.Lock()
	dayEnd:=dayStart.Add(time.Hour*168)
	result := make([]*events.Event, 0)
	for date, event := range s.events {
		if (date.After(dayStart) && date.Before(dayEnd)) || date.Equal(dayStart) || date.Equal(dayEnd) {	
			for _, oneEvent := range event {
				result=append(result, oneEvent)
			}
		}
	}
	defer s.Unlock()
	if result == nil {
		return nil, fmt.Errorf("Нет заметок для этой недели")
	}
	return result, nil
}

func (s *Storage) GetEventForMonth(dayStart time.Time) ([]*events.Event, error) {
	s.Lock()
	dayEnd:=dayStart.Add(time.Hour*5040)
	result := make([]*events.Event, 0)

	for date, event := range s.events {
		if (date.After(dayStart) && date.Before(dayEnd)) || date.Equal(dayStart) || date.Equal(dayEnd) {	
			for _, oneEvent := range event {
				result=append(result, oneEvent)
			}
		}
	}
	defer s.Unlock()
	if result == nil {
		return nil, fmt.Errorf("Нет заметок для этого месяца")
	}
	return result, nil
}

func (s *Storage) Add(title string, date time.Time) (*events.Event, error) {
	s.Lock()
	defer s.Unlock()
	if even, err := events.NewEvent(title, date); err != nil {
		return nil, err
	} else {
		s.events[date]=append(s.events[date], even)
		return even,nil
	}
	
}

func (s *Storage) Delete(day time.Time, title string) (*events.Event,error) {
	s.Lock()
	defer s.Unlock()
	for i, oneEvent := range s.events[day] {
		if oneEvent.Title == title {
			copy(s.events[day][i+1:], s.events[day][i+1:])
			s.events[day][len(s.events[day])-1] = nil
			s.events[day] = s.events[day][:len(s.events[day])-1]
			return oneEvent, nil
		}
	}
	return nil, fmt.Errorf("Нет такой заметки")
	
}

func (s *Storage) Update(day time.Time, title, newTitle string) (*events.Event,error) {
	s.Lock()
	defer s.Unlock()
	for _, oneEvent := range s.events[day] {
		if oneEvent.Title == title {
			oneEvent.Title = newTitle
			return oneEvent, nil
		}
	}
	return nil, fmt.Errorf("Нет таких заметок")
}