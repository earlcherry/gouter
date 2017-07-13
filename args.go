package gouter

type RouteArgs interface {
	GetRoute() string
}

type SimpleRouteArgs struct {
	route string
}

func (a SimpleRouteArgs) GetRoute() string {
	return a.route
}

func (a *SimpleRouteArgs) SetRoute(route string) {
	a.route = route
}
