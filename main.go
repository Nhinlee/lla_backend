package main

import (
	"database/sql"
	"fmt"
	"lla/api"
	"lla/config"
	db "lla/db/sqlc"
	fs "lla/golibs/file_store"

	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	config, error := config.LoadConfig()
	if error != nil {
		log.Fatal("cannot load config: ", error)
	}

	// Connect to database
	conn, err := sql.Open("postgres", config.DBConfig.GetDBConnection())
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	} else {
		fmt.Println("connect to database successfully!")
	}
	store := db.NewStore(conn)

	// Connect to file store (GCS)
	gcsStore, err := fs.NewGCSFileStore()
	if err != nil {
		log.Fatal("cannot connect to file store: ", err)
	}

	runRestfulServer(store, gcsStore)
}

func runRestfulServer(store db.Store, fileStore fs.FileStore) {
	server, err := api.NewServer(store, fileStore)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use a default port if not provided by Heroku
	}

	err = server.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
