package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	*router
	middlewares []Middleware
	startHandler HandlerFunc
}


func NewServer() *Server {
	r := &routeer{make(map[string]map[string]HandlerFunc)}

	s := &Server{router: r}

	s.middlewares = []Middleware{
		logHandler,
		recoverHandler,
		staticHandler,
		parseFormHandler,
		parseJsonBodyHandler
	}
	return s 
	
}
