package main

import (
	"github.com/Online_Song_Libraries/api"
	"github.com/Online_Song_Libraries/api/handler"
	"github.com/Online_Song_Libraries/config"
	_ "github.com/Online_Song_Libraries/docs"
	"github.com/Online_Song_Libraries/storage/postgres"
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
	// host := "library_song_service"
	// port := cfg.HTTPPort
	// if err := router.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
	// 	log.Fatalf("Failed to run server: %v", err)
	// }
}
