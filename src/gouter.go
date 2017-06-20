package gouter

type Gouter struct {
	routes map[string]Gout
}
type Gout struct {
	Controller interface{}
	Method     string
}

func New() *Gouter {
	return &Gouter{
		routes: make(map[string]Gout),
	}
}

func (g *Gouter) Add(rout string, controller interface{}, method string) {
	g.routes[rout] = Gout{
		Controller: controller,
		Method:     method,
	}
}
func (g *Gouter) Run(rout string) {
	gout := g.routes[rout]
	Call(gout.Controller, gout.Method)
}
