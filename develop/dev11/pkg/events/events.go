package events

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {//Тип - событие
	ID uuid.UUID `json:"id"`
	Title string `json:"title"`
	Date time.Time `json:"date"`
}

func NewEvent(title string, date time.Time) (*Event, error) {//Конструктор события
	if title != "" {
		return &Event{
			ID: uuid.New(),
			Title: title,
			Date: date,
		}, nil
	} else {
		return nil, fmt.Errorf("Пустая заметка %s", title)
	}

}