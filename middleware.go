package main

import (
	"log"
	"time"
	"net/http"
)
type Middelware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc)HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		next(c)

		log.Printf("[%s] %q %v\n ",
			c.Request.Method,
			c.Request.URL.String(),
			time.Now().Sub(t))
	}
}


func recoverHandler(next HandlerFunc)HandlerFunc{
	return func (c *Context){
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v",err)
				http.Error(c.ResponseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				
			}
		}()

		next(c)
	}
	
}
