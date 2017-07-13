package gouter

import "errors"

type Gouter struct {
	goutes map[string]Goute
}

func New() *Gouter {
	return &Gouter{
		goutes: make(map[string]Goute),
	}
}

func (g *Gouter) Add(route string, goute Goute) {
	g.goutes[route] = goute
}

func (g *Gouter) Run(args RouteArgs) error {
	var err error

	if goute, ok := g.goutes[args.GetRoute()]; ok {
		err = goute.Run(args)
		if err == nil {
			return nil
		}
	} else {
		err = errors.New("Goute " + args.GetRoute() + " not exist")
	}
	return err
}
