package main

import (
	"fmt"
	"lla/api"
	"lla/config"
	db "lla/db/sqlc"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	config, error := config.LoadConfig()
	if error != nil {
		log.Fatal("cannot load config: ", error)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use a default port if not provided by Heroku
	}

	// Connect to database
	conn, err := sql.Open("postgres", config.DBConfig.GetDBConnection())
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	} else {
		fmt.Println("connect to database successfully!")
	}

	store := db.NewStore(conn)

	runRestfulServer(config, port, store)
}

func runRestfulServer(config *config.Config, port string, store db.Store) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
