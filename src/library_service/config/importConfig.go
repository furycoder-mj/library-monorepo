package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	Port     string
	BasePath string
	UsersURI string
	BooksURI string
}

var Config Configuration

func init() {
	Config.Port = os.Getenv("Port")
	Config.BasePath = os.Getenv("BasePath")
	Config.UsersURI = os.Getenv("UsersServiceURL") + os.Getenv("UsersURI")
	Config.BooksURI = os.Getenv("BooksServiceURL") + os.Getenv("BooksURI")
	fmt.Println(Config)
}
