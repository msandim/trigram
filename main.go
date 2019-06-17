package main

import (
	"github.com/msandim/trigram/server"
	"github.com/msandim/trigram/store"
)

func main() {
	server := server.NewServer(store.NewMapTrigramStore(&store.RandomChooser{}), 8080)
	server.Run()
}
