package main

import (
	"flag"
	"fmt"
	"inv-id-oice/idl"
	"inv-id-oice/web"
	"log"
	"mime"
	"os"
)

func main() {
	var ver = flag.Bool("ver", false, "Prints the current version")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	flag.Parse()

	if *ver {
		fmt.Printf("%s, version: %s", idl.Appname, idl.Buildnr)
		os.Exit(0)
	}
	serv := web.NewServiceRunner(*configfile)
	if err := serv.RunService(); err != nil {
		log.Fatal("[MAIN] unable to run the service ", err)
	}
}

func init() {
	_ = mime.AddExtensionType(".js", "text/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
	_ = mime.AddExtensionType(".mjs", "text/javascript")
	log.Printf("Init App %s with Mime override", idl.Appname)
}
