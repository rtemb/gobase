package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
)

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Required env variable PORT is not set")
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", root)
	r.GET("/info", info)
	r.GET("/healthz", healthz)

	log.Info("Service started up at port: " + port)
	r.Listen(":" + port)
}
