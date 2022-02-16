package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/virtsevda/StandartWebServer/storage"
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

//Пытаемся скофнигурировать наше хранилище
func (a *API) configureStorageField() error{
	storage := storage.New(a.config.Storage)

	if err:=storage.Open(); err!=nil{
		return err
	}
	a.storage = storage
	return nil
}
