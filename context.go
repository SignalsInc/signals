// Smoke Signals
// We give you the fire, you provide the signals.
package main

import (
	"log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/FoundationDB/fdb-go/fdb"
)

// If Go had generics I could generalize this ContextHandler concept without
// resorting to casting the Context type.
type Context struct {
	router *mux.Router
	db fdb.Database
}
type ContextFunc func(*Context, http.ResponseWriter, *http.Request)
type ContextHandler struct {
    ctx *Context
    fn ContextFunc
}
func (handler ContextHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler.fn(handler.ctx, writer, request)
}

func NewContext() *Context {
	// Router
	router := mux.NewRouter()

	// Database
	err := fdb.APIVersion(100)
	if err != nil {
		log.Fatal(err)
	}

	db, err := fdb.OpenDefault()
	if err != nil {
		log.Fatal(err)
	}

	return &Context{router, db}
}

