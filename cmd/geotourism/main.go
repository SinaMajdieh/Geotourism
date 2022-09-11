package main

import (
	"github.com/SinaMajdieh/Geotourism/internal"
	"github.com/SinaMajdieh/Geotourism/pkg/cli"
	"github.com/SinaMajdieh/Geotourism/pkg/file_pkg"
	"log"
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
