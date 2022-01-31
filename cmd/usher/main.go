package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	"github.com/nozgurozturk/usher/internal/server"
)

func main() {
	// Read env variables
	addr := getEnv("PORT", "8080")
	dbURL := getEnv("DATABASE_URL", "file:usher.db?cache=shared&mode=memory&_fk=1")
	dbDriver := getEnv("DATABASE_DRIVER", "sqlite3")

	// Override if arguments are existed
	port := flag.String("port", addr, "Server PORT")
	databaseURL := flag.String("database", dbURL, "Database URL")
	databaseDriver := flag.String("driver", dbDriver, "Database Driver")

	flag.Parse()

	// Connect DB
	db, err := ent.Open(*databaseDriver, *databaseURL)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Create tables
	if err := db.Schema.Create(ctx, schema.WithDropColumn(true), schema.WithDropIndex(true)); err != nil {
		panic(err)
	}

	// Create Application
	app := application.NewApplication(db)

	// Create HTTP Server
	srv := server.NewHttpServer(app)

	srvErr := make(chan error, 1)
	// Run HTTP Server
	go func() {
		log.Println("HTTP Server is started on:" + *port)
		srvErr <- http.ListenAndServe(":"+*port, srv)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-srvErr:
		log.Println("HTTP Server is stopped")
	case <-quit:
		log.Println("HTTP Server is stopped")
	}

	// Close DB
	if err := db.Close(); err != nil {
		panic(err)
	}
}

//getEnv parse env variables if exist, otherwise it returns default value
func getEnv(env, defaultValue string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultValue
	}
	return e
}
