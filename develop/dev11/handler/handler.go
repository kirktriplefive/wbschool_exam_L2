package handler

import (
	"net/http"

	//"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/service"
)

type HandlerInterface interface {
    GetEventForDay(w http.ResponseWriter, r *http.Request) 
	GetEventForWeek(w http.ResponseWriter, r *http.Request) 
	GetEventForMonth(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
    HandlerInterface
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{&EventHandler{*service}}
}