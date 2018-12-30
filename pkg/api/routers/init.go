package routers

import (
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/api/controllers"
	"github.com/alamin-mahamud/golang-jwt-authentication-api-sample/pkg/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() http.Handler{
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}

func SetHelloRoutes(r *mux.Router) *mux.Router{
	r.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
			)).Methods("GET")
	return router
}
func SetAuthenticationRoutes(r *mux.Router) *mux.Router{
	r.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	r.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
			)).Methods("GET")

}