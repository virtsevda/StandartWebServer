package api

import "github.com/sirupsen/logrus"

//Base API server instace
type API struct {
	//UNEXPORTED field
	config *Config
	logger *logrus.Logger
}

//Api constructor: buil base API instace
func New(config *Config) *API{
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

//start http server/configure loggers, router, database connection, etc
func (api *API) Start() error{
	//Try to configure logger
	if err:=api.configureLoggerField(); err!=nil{
		return err
	}
	
	api.logger.Info("Starting api server at port:",api.config.BindAddr)

	return nil
}