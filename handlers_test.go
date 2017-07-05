package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
)

// buffer is a special variable to save log messages
var buffer bytes.Buffer

func init() {
	log.Out = &buffer
	log.Formatter = &logrus.JSONFormatter{}
}

// TestLogger checks if logger handler works correctly
func TestLogger(t *testing.T) {
	r := router.New()
	r.Logger = logger

	ts := httptest.NewServer(r)
	defer ts.Close()

	_, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	formated := struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Time  string `json:"time"`
	}{}
	err = json.NewDecoder(&buffer).Decode(&formated)
	if err != nil {
		t.Fatal(err)
	}

	msgParts := strings.Split(formated.Msg, " ")
	if len(msgParts) != 3 {
		t.Fatalf("Wrong message was logged: %s", formated.Msg)
	}
}