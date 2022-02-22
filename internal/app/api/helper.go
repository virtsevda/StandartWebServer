package api

import (
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/virtsevda/StandartWebServer/storage"
)

var (
	prefix string ="/api/v1"
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


	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles/{id}", a.UpdateArticleById).Methods("PUT")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")

	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
	
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


