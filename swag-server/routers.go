/*
 * Sample REST server
 *
 * TODO
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"DueYearMonthDayGet",
		strings.ToUpper("Get"),
		"/due/{year}/{month}/{day}",
		DueYearMonthDayGet,
	},

	Route{
		"TagTagnameGet",
		strings.ToUpper("Get"),
		"/tag/{tagname}",
		TagTagnameGet,
	},

	Route{
		"TaskGet",
		strings.ToUpper("Get"),
		"/task",
		TaskGet,
	},

	Route{
		"TaskIdDelete",
		strings.ToUpper("Delete"),
		"/task/{id}",
		TaskIdDelete,
	},

	Route{
		"TaskIdGet",
		strings.ToUpper("Get"),
		"/task/{id}",
		TaskIdGet,
	},

	Route{
		"TaskPost",
		strings.ToUpper("Post"),
		"/task",
		TaskPost,
	},
}

func InitRouteSwagger(port string) {
	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
