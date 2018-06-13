package main

import (
	"fmt"

)

func main() {
	s := NewServer()

	s.HandlerFunc("GET", "/", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "welcome!")
	})

	s.HandlerFunc("GET", "/about", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "about")
	})

	s.HandlerFunc("GET", "/users/:id", logHandler(recoverHandler(func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n", c.Params["id"])
	})))

	s.HandlerFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "retrieve  %v's address %v\n", c.Params["user_id"], c.Params["address_id"])
	})

	s.HandlerFunc("POST", "/users", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})))))

	s.Run(":8080")
}
