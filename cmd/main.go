package main

import (
	"github.com/bhkfazano/simple-kv/pkg/server"
	"github.com/bhkfazano/simple-kv/pkg/store"
)

func main() {
	var simpleKV *store.SimpleKV[string, string] = store.NewSimpleKV[string, string]()
	var server *server.Server[string, string] = server.NewServer[string, string](simpleKV, ":8080")

	server.ListenAndServe()
}
