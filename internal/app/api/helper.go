package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Пытаем сконфигурировать api instance
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)

	if err!=nil {
		return err
	}

	a.logger.SetLevel(log_level)
	return nil
}

//Пытаемся сконфигурировать маршрутизатор

func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello! API!"))
	})

}