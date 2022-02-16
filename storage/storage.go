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