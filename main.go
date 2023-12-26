package main

import (
	"flag"
	"fmt"

	gin_router "github.com/victor8titov/apitasks/routers/gin"
	gorilla_routers "github.com/victor8titov/apitasks/routers/gorilla"
	simple_router "github.com/victor8titov/apitasks/routers/simple"
	swagger "github.com/victor8titov/apitasks/swag-server"
)

func main() {
	way := flag.String("way", "simple", "Chose way making rest api. You can use simple || gorilla || gin || swagger")
	port := flag.String("port", "4112", "port for running server")
	flag.Parse()
	fmt.Println("way:", *way)
	fmt.Println("Port:", *port)

	if *way == "simple" {
		simple_router.InitSimpleWayRouter(*port)
	}

	if *way == "gorilla" {
		gorilla_routers.InitGorillaRouter(*port)
	}

	if *way == "gin" {
		gin_router.InitRouterGin(*port)
	}

	if *way == "swagger" {
		swagger.InitRouteSwagger(*port)
	}
}
