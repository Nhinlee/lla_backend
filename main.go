package main

import (
	"context"
	"fmt"
	"lla/api"
	"lla/config"
	db "lla/db/sqlc"
	fs "lla/golibs/file_store"
	"lla/golibs/vision_ai"

	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, error := config.LoadConfig()
	if error != nil {
		log.Fatal("cannot load config: ", error)
	}

	// Create a connection pool
	configDB := config.DBConfig.GetDBConnection()
	poolConfig, err := pgxpool.ParseConfig(configDB)
	if err != nil {
		log.Fatal("cannot parse connection pool config: ", err)
	}
	poolConfig.MaxConns = 10

	// Connect to database
	pool, err := pgxpool.New(context.Background(), poolConfig.ConnString())
	if err != nil {
		log.Fatal("cannot create connection pool: ", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)

	// Connect to file store (GCS)
	gcsStore, err := fs.NewGCSFileStore()
	if err != nil {
		log.Fatal("cannot connect to file store: ", err)
	}

	// Setup authentication for Vision API
	visionAI := vision_ai.NewVisionAI()

	runRestfulServer(store, gcsStore, visionAI)
}

func runRestfulServer(store db.Store, fileStore fs.FileStore, visionAI *vision_ai.VisionAI) {
	server, err := api.NewServer(store, fileStore, visionAI)
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
