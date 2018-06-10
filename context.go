package main

import (
	"fmt"
	"net/http"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request *http.Request
}

type HanderFunc func(*Context)
