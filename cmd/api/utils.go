package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]interface{}

// function to read incoming JSON
func readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    // set max amount of Bytes to be read 
    maxBytes := 1024 * 1024 // one megabyte
    r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

    dec := json.NewDecoder(r.Body)

    dec.DisallowUnknownFields()

    // decode the JSON object
    err := dec.Decode(dst)
    if err != nil {
        return err
    }

    // decode the object again into an empty struct.
    // This is too check if there are two JSON object present in the payload
    err = dec.Decode(&struct{}{})
    if err != io.EOF {
        return errors.New("body must only contain a single JSON value")
    }

    return nil
}

// function to write JSON
func writeJSON(w http.ResponseWriter, status int, data envelope, header http.Header) error {
    // Encode the data to JSON 
    js, err := json.Marshal(data)
    if err != nil {
        return err
    }

    // append new line so its easier to interpret
    js = append(js, '\n')

    for key, value := range header {
        w.Header()[key] = value
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)

    return nil
}
