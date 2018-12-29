package main

import (
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/routers"
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
