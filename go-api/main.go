package main

import (
	"flag"
	"os"

	"github.com/hikvineh/go-api/cmd/system/app"
	DB "github.com/hikvineh/go-api/cmd/system/db"
	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8081", "Assigning the port that the server should listen to.")

	flag.Parse()
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	// log.Println("Hello World!")
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
