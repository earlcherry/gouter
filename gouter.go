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

func (g *Gouter) Run(args Args) error {
	var err error
	if goute, ok := g.goutes[args.Command]; ok {
		err = goute.Run(args)
		if err == nil {
			return nil
		}
	} else {
		err = errors.New("Goute " + args.Command + " not exist")
	}
	return err
}
