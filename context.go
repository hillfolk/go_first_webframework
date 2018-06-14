package main

import (
	"net/http"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) RenderJson(v interface{}){

	c.ResponseWriter.WriteHeader(http.StatusOK)
	c.ResponseWriter.Header().set("Content-Type","application/json;charset=utf8")

	if err := json.NewEncoder(c.ResponseWriter).Encode(v);err != nil {
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderXml(v interface{}){
	c.ResponseWriter.WriterHeader(http.StatusOK)

	c.ResponseWriter.Header().set("Content-Type","application/xml;charset=utf8")

	if err := json.NewEncoder(c.ResponseWriter).Encode(v);err != nil {
		c.RenderErr(http.StatusInternalServerError, err)
	}
}


func (c *Context) RenderErr(code int, err error){
	if err != nil {
		if code > 0 {
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr),defaultErr)
		}
	}
}



type HandlerFunc func(*Context)
