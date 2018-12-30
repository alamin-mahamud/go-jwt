package main

import (
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/api/routers"
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/settings"
	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	router := routers.Init()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run()
}