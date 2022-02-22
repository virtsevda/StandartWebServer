package storage

import (
	"fmt"
	"log"

	"github.com/virtsevda/StandartWebServer/internal/app/models"
)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle="articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article,error){
	query:=fmt.Sprintf("INSERT INTO %s (title,author,content) VALUES ($1,$2,$3) REURNING id",tableArticle)
	if err:= ar.storage.db.QueryRow(query,&a.Title,&a.Author,&a.Content).Scan(&a.ID);err!=nil{
		return nil,err
	}
	return a,nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article,error){
	article,ok,err:=ar.FindById(id)
	if err!=nil{
		return nil,err
	}
	if ok {
		query:=fmt.Sprintf("DELETE FROM %s WHERE id=$1",tableArticle)
		_,err:=ar.storage.db.Exec(query,id)
		if err!=nil{
			return nil,err
		}
	}
	return article,nil
}

func (ar *ArticleRepository) FindById(id int) (*models.Article,bool,error){
	articles, err:=ar.SelectAll()
	var founded bool
	if err!=nil{
		return nil,founded,err
	}
	var articleFinded *models.Article
	for _,a:= range articles{
		if a.ID == id{
			articleFinded =a
			founded=true
			break
		}
	}
	return articleFinded,founded,nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article,error){
	query:=fmt.Sprintf("SELECT *FROM %s",tableArticle)
	fmt.Println(query)
	rows,err:=ar.storage.db.Query(query)
	if err!=nil{
		return nil,err
	}
	articles := make([]*models.Article,0)
	for rows.Next(){
		a:=models.Article{}
		err:= rows.Scan(&a.ID,&a.Title,&a.Author,&a.Content)
		if err!=nil{
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	defer rows.Close()
	return articles,nil
}

func (ar *ArticleRepository) UpdateById(a *models.Article) (*models.Article,error){
	query:=fmt.Sprintf("UPDATE %s SET title=$1, author=$2,content=$3 WHERE id=$4 REURNING id",tableArticle)
	if err:= ar.storage.db.QueryRow(query,&a.Title,&a.Author,&a.Content).Scan(&a.ID);err!=nil{
		return nil,err
	}
	return a,nil
}