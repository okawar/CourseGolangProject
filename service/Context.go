package service

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type errorMessage struct {
	Mes string
}

func (c *Context) sendError(status int, mes string) {
	em := errorMessage{Mes: mes}
	marsh, _ := json.Marshal(&em)
	c.Response.Write(marsh)
	c.Response.WriteHeader(status)
}

func (c *Context) sendAnswer(data interface{}) {
	c.Response.Header().Set("Content-Type", "text/json")
	marsh, _ := json.Marshal(data)
	c.Response.Write(marsh)
}

/*
	func (c *Context) redirect(destination string) {
		c.Response.Header().Set("Location", destination)
		c.Response.WriteHeader(303)
	}
*/
func GetData[T any](c *Context) (ent T, err error) {
	err = json.NewDecoder(c.Request.Body).Decode(&ent)
	return
}
