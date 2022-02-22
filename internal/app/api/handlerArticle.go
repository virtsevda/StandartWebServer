package api

import (
	"encoding/json"
	"net/http"
)


func (api *API)GetAllArticles(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Handler: GetAllArticles, url:/api/v1/articles, method:GET")
	articles,err := api.storage.Article().SelectAll()
	if err!=nil{
		api.logger.Info("Error!!!Handler: GetAllArticles, url:/api/v1/articles, method:GET")
		msg:=Message{
			StatusCode: 501,
			Message:"We have some troubles to acessing database. Try again later",
			IsError:true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(articles)

}

func (api *API)GetArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Handler: GetArticleById, url:/api/v1/articles/{id}, method:GET")
}

func (api *API)DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Handler: DeleteArticleById, url:/api/v1/articles/{id}, method:DELETE")
}

func (api *API)UpdateArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Handler: UpdateArticleById, url:/api/v1/articles/{id}, method:PUT")
}

func (api *API)PostArticle(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Handler: PostArticle, url:/api/v1/articles, method:POST")
}
