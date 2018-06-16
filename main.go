package main

import (
	"fmt"
	"time"

)

type User struct {
	Id        string
	AddressId string
}


func main() {
	s := NewServer()

	s.HandlerFunc("GET", "/", func(c *Context) {
		c.RenderTemplate("/public/index.html",map[string]interface{}{"time":time.Now()})
	})

	s.HandlerFunc("GET", "/about", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "about")
	})

	s.HandlerFunc("GET", "/users/:id", logHandler(recoverHandler(func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		u := User{Id:c.Params["id"].(string)}
		c.RenderXml(u)
	})))

	s.HandlerFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "retrieve  %v's address %v\n", c.Params["user_id"], c.Params["address_id"])
		u := User{Id:c.Params["user_id"].(string),AddressId:c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	s.HandlerFunc("POST", "/users", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})))))

	s.Run(":8080")
}
