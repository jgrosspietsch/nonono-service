package main

import (
	"flag"
	"fmt"

	"github.com/jgrosspietsch/nonono-service/api"
)

var listenOn string
var production bool

func init() {
	flag.StringVar(&listenOn, "listen", ":8888", "string passed into gin to listen on")
	flag.BoolVar(&production, "prod", false, "are we running in production?")
}

func main() {
	flag.Parse()
	fmt.Println("listenOn", listenOn)
	fmt.Println("production", production)

	router := api.SetupInitialRouter(listenOn, production)
	api.SetupRoutes(router)
	
	router.Run(listenOn)
}
