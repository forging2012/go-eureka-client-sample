package service

import "net/http"

type Router struct {
	PathPrefix string
	Routes     Routes
}

type Routers []Router

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routers = Routers{
	Router{
		PathPrefix: "/",
		Routes: Routes{
			Route{
				Name:        "Index",
				Method:      "GET",
				Pattern:     "/",
				HandlerFunc: Index,
			},
			Route{
				Name:        "Info",
				Method:      "GET",
				Pattern:     "/info",
				HandlerFunc: Info,
			},
			Route{
				Name:        "Health",
				Method:      "GET",
				Pattern:     "/health",
				HandlerFunc: Health,
			},
			Route{
				Name:        "VendorShow",
				Method:      "GET",
				Pattern:     "/vendors/{productId}",
				HandlerFunc: VendorShow,
			},
		},
	},
}
