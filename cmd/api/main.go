package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/virtsevda/StandartWebServer/internal/app/api"
)

var (
	configPath string
	format string
)

func init() {
	//
	flag.StringVar(&format,"format","toml","Format config file .toml or .env")
	flag.StringVar(&configPath,"path","configs/api.toml","Path to config file in toml or .env format")
}

func main() {
	log.Println("It works")
	flag.Parse()

	//server instance initialization
	config := api.NewConfig()
	if format == "toml"{
		_,err := toml.DecodeFile(configPath,&config)
		if err !=nil{
			log.Println("Can not find configs file. Using default values:",err)
		}
		log.Println("Load toml")
	}

	if format == "env"{
		err := godotenv.Load(configPath)
		if err != nil {
			log.Println("Can not find configs file. Using default values:",err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
		config.Storage.DatabaseURI = os.Getenv("databse_uri")
	}


	server := api.New(config)

	if err := server.Start(); err !=nil{
		log.Fatal(err)
	}

	//log.Fatal(server.Start())

}