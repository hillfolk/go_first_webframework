package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandlerFunc("GET", "/", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "welcome!")
	})

	r.HandlerFunc("GET", "/about", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "about")
	})

	r.HandlerFunc("GET", "/users/:id",logHandler(recoverHandler(func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n",c.Params["id"])
	})))

	r.HandlerFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "retrieve  %v's address %v\n",c.Params["user_id"],c.Params["address_id"])
	})
	r.HandlerFunc("GET", "/users", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "create user\n")
	})
	

	
	
	http.ListenAndServe(":8080", r)
}
