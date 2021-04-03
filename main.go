// Package main
package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortener/api"
	"urlshortener/redirect"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Init
func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	fmt.Println("Url Shortener App!!!")
	// Check if env variables are set if not set it to default values
	if os.Getenv("LISTEN_PORT") == "" {
		os.Setenv("LISTEN_PORT", "8080")
	}
	if os.Getenv("LOGFILE_PATH") == "" {
		os.Setenv("LOGFILE_PATH", "./urlshortener_log.out")
	}
	// Get the listing port number and logfile path from os env variables
	logrus.Info("LISTEN_PORT  : " + os.Getenv("LISTEN_PORT"))
	logrus.Info("LOGFILE PATH : " + os.Getenv("LOGFILE_PATH"))
	port, logpath := os.Getenv("LISTEN_PORT"), os.Getenv("LOGFILE_PATH")
	f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		logrus.SetOutput(f)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", api.GetVersion).Methods("GET")
	router.HandleFunc("/shrink", api.GetShortenedURL).Methods("GET")
	router.HandleFunc("/shrink", api.CreateShortenedURL).Methods("POST")
	router.HandleFunc("/redirect", redirect.RedirectURL).Methods("GET")
	logrus.Info("Starting server and listening on port " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		logrus.Error(err)
	}
}
