package frame

import (
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
	"main/service"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range service.RoutesArray {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		fmt.Printf(
			"Servlet name is %s , parse to url : %s, http type is %s\n",
			route.Name,
			route.Pattern,
			route.Method,
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
