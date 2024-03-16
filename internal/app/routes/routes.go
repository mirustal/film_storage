package routes

import "net/http"


type Route struct {
	Method	  string
	Pattern	 string
	HandlerFunc http.HandlerFunc
}

type Routes map[string]Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}


func NewRouter(routers ...Router) *http.ServeMux {
	router := &http.DefaultServeMux

	for _, api := range routers {
		for name, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = Logger(handler, name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(name).
				Handler(handler)
		}
	}

	return router
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r.mux)
}