package main

import (
	"github.com/Online_Song_Libraries/api"
	"github.com/Online_Song_Libraries/api/handler"
	"github.com/Online_Song_Libraries/config"
	"github.com/Online_Song_Libraries/storage/postgres"
	_ "github.com/Online_Song_Libraries/docs"
)	

func main() {
	cfg := config.Load()
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	handler := handler.NewHandler(db)

	router := api.NewGin(handler)

	router.Run(cfg.HTTPPort)

}
