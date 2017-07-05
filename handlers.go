package main

import (
	"fmt"

	"github.com/rtemb/gobase/appinfo"
	"github.com/takama/router"
	"net/http"
)

// root test endpoint
func root(c *router.Control) {
	fmt.Fprintf(c.Writer, "Processing URL %s...", c.Request.URL.Path)
}

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}

// Readiness probe for Kubernetes
func info(c *router.Control) {
	appinfo.Info(c, appinfo.RELEASE, appinfo.REPO, appinfo.COMMIT)
}

// Liveness probe for Kubernetes
func healthz(c *router.Control) {
	c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
}
