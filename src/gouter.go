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

func (g *Gouter) Add(route string, Goute Goute) {
	g.goutes[route] = Goute
}

func (g *Gouter) Run(route string) error {
	if goute, ok := g.goutes[route]; ok {
		goute.Run()
		return nil
	} else {
		return errors.New("Goute " + route + " not exist")
	}
}
