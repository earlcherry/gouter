package gouter

type Goute interface {
	Run(args RouteArgs) error
}
