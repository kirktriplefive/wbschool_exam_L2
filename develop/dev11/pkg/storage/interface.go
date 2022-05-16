package storage

import (
	"sync"
	"time"

	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
)


type StorageInterface interface {
	GetEventForDay(day time.Time) ([]*events.Event, error)
	GetEventForWeek(dayStart time.Time) ([]*events.Event, error)
	GetEventForMonth(dayStart time.Time) ([]*events.Event, error)
	Add(title string, date time.Time) (*events.Event, error)
	Delete(day time.Time, title string) (*events.Event,error)
	Update(day time.Time, title, newTitle string) (*events.Event,error)
}

type Storage struct {//тип для сохранения ивентов в памяти. Решено организовать в виде мапы с ключом - датой и значением - слайсом ивентов
	sync.Mutex
	events map[time.Time][]*events.Event
}

func NewStorage() *Storage {//конструктор для сохранения
	eventStorage := make(map[time.Time][]*events.Event)
	return &Storage{events: eventStorage}
}