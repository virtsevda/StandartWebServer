package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Base API server instace
type API struct {
	//UNEXPORTED field
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//Api constructor: buil base API instace
func New(config *Config) *API{
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//start http server/configure loggers, router, database connection, etc
func (api *API) Start() error{
	//Try to configure logger
	if err:=api.configureLoggerField(); err!=nil{
		return err
	}
	
	api.logger.Info("Starting api server at port:",api.config.BindAddr)

	//Конфигирурируем роутер
	api.configureRouterField()

	//

	return http.ListenAndServe(api.config.BindAddr,api.router)
}