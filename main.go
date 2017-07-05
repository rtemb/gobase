package main

import (
	"os"

	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/rtemb/rprint/info"
	"github.com/rtemb/rprint/version"
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

	// Readiness and liveness probes for Kubernetes
	r.GET("/info", func(c *router.Control) {
		info.Info(c, version.RELEASE, version.REPO, version.COMMIT)
	})
	r.GET("/healthz", func(c *router.Control) {
		c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
	})

	log.Info("Service started up at port: " + port)
	r.Listen(":" + port)
}
