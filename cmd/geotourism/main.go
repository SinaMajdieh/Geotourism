package main

import (
	"log"

	"Geotourism/internal"
	"Geotourism/pkg/cli"
	"Geotourism/pkg/file_pkg"
)

func main() {
	internal.DeclarePages()
	internal.Load_Data()
	var server internal.Server
	server = internal.NewHttpServer("", "")
	err := file_pkg.ReadJson("configs/server.json", server)
	if nil != err {
		log.Fatal("config file : ", err)
	}
	go cli.Cmd()
	server.SetupServer()

}
