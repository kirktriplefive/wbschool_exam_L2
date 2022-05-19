package main

import (
	"net/http"
	"os"

	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/handler"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/middleware"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/service"
	"github.com/kirktriplefive/wbschool_exam_L2/develop/dev11/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	storage := storage.NewStorage()// инициализируем структуру Storage
	service := service.NewService(storage)// инициализируем структуру Service
	handler := handler.NewHandler(service)// инициализируем структуру Handler
	mux := initRouter(handler) // инициализируем роутер
	errChan := make(chan error) // канал для получения ошибок работы сервера
	if err := initConfig(); err != nil { // инициализируем конфиг для того, чтобы узнать порт
		logrus.Fatalf("error init config: %v", err)
	}
	go func() {
		errChan <- http.ListenAndServe(viper.GetString("port"), mux) // запускаем сервер
	}()
	var err error
	select {
	case err = <-errChan: // проверяем нет ли ошибок в работе сервера
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	}

}

func initRouter(h *handler.Handler) *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/create_event", middleware.Logging(http.HandlerFunc(h.Add))) // хэндлеры
	mux.HandleFunc("/update_event", middleware.Logging(http.HandlerFunc(h.Update)))
	mux.HandleFunc("/delete_event", middleware.Logging(http.HandlerFunc(h.Delete)))
	mux.HandleFunc("/events_for_day", middleware.Logging(http.HandlerFunc(h.GetEventForDay)))
	mux.HandleFunc("/events_for_week", middleware.Logging(http.HandlerFunc(h.GetEventForWeek)))
	mux.HandleFunc("/events_for_month", middleware.Logging(http.HandlerFunc(h.GetEventForMonth)))

	return mux
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

type server struct {
	httpServer *http.Server
}



