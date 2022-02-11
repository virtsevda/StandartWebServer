package api

//Base API server instace
type API struct {
	//UNEXPORTED field
	config *Config
}

//Api constructor: buil base API instace
func New(config *Config) *API{
	return &API{
		config: config,
	}
}

//start http server/configure loggers, router, database connection, etc
func (api *API) Start() error{
	return nil
}