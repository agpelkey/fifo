package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// function to start the server
func (app *application) serve() error {
    
    // declare server settings
    srv := http.Server{
        Addr: fmt.Sprintf(":%d", app.config.port),
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }
    
    // create channel to receive any errors returned by our graceful shutdown
    shutdownError := make(chan error)

    // start a background go routine to run for the lifetime of the application
    go func() {
        // create another channel to catch all signals coming from the server
        quit := make(chan os.Signal, 1)

        // use signal.Notify to listen for incoming signals and send them to the quit channel
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

        // Read the signals from the quit channel into a variable
        s := <- quit

        // log out the signal message from the channel
        log.Println(s)

        // create server context
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        // call Shutdown(). Relay the return value from the function to the shut down channel
        shutdownError <- srv.Shutdown(ctx)
    }()

    fmt.Println("starting server")

    // start the server. Note that if shutdown() is executed succesfully then ListenAndServe
    // will return an http.ErrServerClosed. So, we check for any error that is *not* this
    err := srv.ListenAndServe()
    if !errors.Is(err, http.ErrServerClosed) {
        return err
    }

    // otherwise, wait for the return value from our call to shutdown().
    // If return value is an error, then we now know our graceful shutdown
    // encountered an issue.
    err = <- shutdownError
    if err != nil {
        return err
    }

    log.Println("stopping server")

    return nil
}
