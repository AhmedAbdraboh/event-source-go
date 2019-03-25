package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Route struct {
	Path string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
}


func SetupRoutes() *chi.Mux{
	var routes []Route
	routes = append(routes, accountRoutes...)

	r := chi.NewRouter()
	for _, route := range routes{
		r.MethodFunc(route.Method, route.Path, route.Handler)
	}
	return r
}

