package service

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	muxRouter := mux.NewRouter().StrictSlash(true)
	for _, router := range routers {
		muxSubRouter := muxRouter.PathPrefix(router.PathPrefix).Subrouter()
		for _, route := range router.Routes {
			handler := Logger(route.HandlerFunc, route.Name)
			muxSubRouter.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	return muxRouter
}
