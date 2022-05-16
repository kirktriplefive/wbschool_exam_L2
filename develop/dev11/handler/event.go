package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/events"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/service"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	service service.ServiceInterface
}

type Event struct {
	Title string `json:"title"`
	Date string `json:"date"`
}

type UpdateEvent struct {
	Title string `json:"title"`
	Date string `json:"date"`
	NewTitle string `json:"new_title"`
}


func NewEventHandler(service service.ServiceInterface) *EventHandler {
	return &EventHandler{
		service: service, 
	}
}

func (h *EventHandler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("GET params were:", r.URL.Query())
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		jsonStr := r.URL.Query().Get("date")
		e, err := h.service.GetEventForDay(jsonStr)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendSliceResponse(w, http.StatusOK, e)
	}
	
	
}

func (h *EventHandler) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		jsonStr := r.URL.Query().Get("date")
		e, err := h.service.GetEventForWeek(jsonStr)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendSliceResponse(w, http.StatusOK, e)
	}
}

func (h *EventHandler) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен GET").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		jsonStr := r.URL.Query().Get("date")
		e, err := h.service.GetEventForMonth(jsonStr)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendSliceResponse(w, http.StatusOK, e)
	}
}


func (h *EventHandler) Delete(w http.ResponseWriter, r *http.Request)  {
	var event Event
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusBadRequest,
		})
		logrus.Println("Error decoding", err.Error(), event)
		return
		}
		e, err := h.service.Delete(event.Date, event.Title)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendResponse(w, http.StatusOK, e)
	}
	
}

func (h *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var event UpdateEvent
	// if r.Body == nil {
	// 	http.Error(w, "Please send a request body", 400)
	// 	logrus.Println("nil body")
	// 	return
	// }
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusBadRequest,
		})
		logrus.Println("Error decoding", err.Error(), event)
		return
		}
		e, err := h.service.Update(event.Date, event.Title, event.NewTitle)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendResponse(w, http.StatusOK, e)
	}
}

func (h *EventHandler) Add(w http.ResponseWriter, r *http.Request) {
	var event Event
	if r.Method != "POST" {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          fmt.Errorf("Неверный метод запроса, нужен POST").Error(),
			HTTPStatusCode: http.StatusMethodNotAllowed,
		})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusBadRequest,
		})
		logrus.Println("Error decoding", err.Error(), event)
		return
		}
		e, err := h.service.Add(event.Title, event.Date)
		if err != nil {
		sendErrorResponse(w, r, &ErrorModel{
			Error:          err.Error(),
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		logrus.Println(err)
		return
		}
		sendResponse(w, http.StatusOK, e)
	}
}

type ErrorModel struct {
	Error string
	HTTPStatusCode int
}

func sendErrorResponse(w http.ResponseWriter, r *http.Request, e *ErrorModel) {
	w.WriteHeader(e.HTTPStatusCode)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(e)
	w.Write(reqBodyBytes.Bytes())

}

func sendResponse(w http.ResponseWriter, s int, event *events.Event) {
	w.WriteHeader(s)
	data, err := json.MarshalIndent(event, "", "    ")
    if err != nil {
        logrus.Fatal(err)
    }
	w.Write(data)
}

func sendSliceResponse(w http.ResponseWriter, s int, event []*events.Event) {
	w.WriteHeader(s)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(event)
	data, err := json.MarshalIndent(event, "", "    ")
    if err != nil {
        logrus.Fatal(err)
    }
	w.Write(data)
}



