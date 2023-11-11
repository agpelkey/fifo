package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/agpelkey/food"
	"github.com/agpelkey/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)


type config struct {
    port int
    env string
}

type application struct {
    config config
    // add future db interface connections here
    ProteinStore food.ProteinService
}

func main() {

    // create config
    var cfg config

    // pass in server arguments
    flag.IntVar(&cfg.port, "port", 8080, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()


    // add db connection
    db, err := pgxpool.New(context.Background(), os.Getenv("FIFO_DSN"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to create database connection")
        os.Exit(1)
    }

    fmt.Println("database connection established")

    app := &application{
        config: cfg,
        // add futue application configs here
        ProteinStore: postgres.NewProteinStore(db),
    }

    // start server
    err = app.serve()
    if err != nil {
        log.Fatal(err)
    }
}
