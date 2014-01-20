// Smoke Signals
// We give you the fire, you provide the signals.
package main

import (
	"log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/SignalsInc/fdb-go/fdb"
)

func PostSignalHandler(ctx *Context, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key, value := vars["key"], vars["value"]
	err := ctx.db.Set(fdb.Key(key), []byte(value))
	if err != nil {
		log.Printf("Key value could not be set: %s -> %s", key, value)
		writer.WriteHeader(400)
		writer.Write([]byte("Failed"))
		return
	}

	writer.WriteHeader(200)
	writer.Write([]byte("Success"))
}

func GetSignalHandler(ctx *Context, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value, err := ctx.db.Get(fdb.Key(key))
	if err != nil {
		log.Printf("Value could not be retrieved for key: %s", key)
		writer.WriteHeader(404)
		return
	}

	writer.WriteHeader(200)
	writer.Write(value)
}

func main() {
	ctx := NewContext()
	ctx.router.Methods("GET").Path("/channels/{key}").Handler(ContextHandler{ctx, GetSignalHandler})
	ctx.router.Methods("POST").Path("/channels/{key}/{value}").Handler(ContextHandler{ctx, PostSignalHandler})
	ctx.router.Methods("GET").Path("/static").Handler(http.FileServer(http.Dir("/Users/jhuffaker/Projects/go/src/github.com/SignalsInc/ssignal/static")))
	http.ListenAndServe(":8080", ctx.router)
}

