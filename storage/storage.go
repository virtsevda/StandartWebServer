package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance

type Storage struct{
	config *Config
	db *sql.DB
	userRepository *UserRepository
	articleRepository *ArticleRepository

}

func New(config *Config) *Storage {
	return &Storage{
		config:config,
	}
}


func (storage *Storage) Open() error{
	db, err := sql.Open("postgres", storage.config.DatabaseURI)

	if err != nil {		
		return err
	}

	if err:= db.Ping(); err!=nil {
		return err
	}
	
	storage.db = db
	log.Println("Database connection successfully")
	return nil
}

func (storage *Storage) Close() {
	storage.db.Close()
}




//public Repo for User
func (storage *Storage) User() *UserRepository{
	if storage.userRepository !=nil{
		return storage.userRepository
	}
	storage.userRepository = &UserRepository{
		storage: storage,
	}
	return nil
}

//public Repo for Article
func (storage *Storage) Article() *ArticleRepository{
	if storage.articleRepository !=nil{
		return storage.articleRepository
	}
	storage.articleRepository = &ArticleRepository{
		storage: storage,
	}
	return nil
}