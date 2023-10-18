package main

import "flag"


type config struct {
    port int
    env string
}

type application struct {
    config config

    // add future db interface connections here
}

func main() {

    // create config
    var cfg config

    // pass in server arguments
    flag.IntVar(&cfg.port, "port", 8080, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()


    // add db connection

    app := &application{
        config: cfg,
        // add futue application configs here
    }

    // start server

}
