package goreact

type buildData struct {
	engine *Engine
	el     *NodeData
}

type Component interface {
	getName() string

	build(data buildData) Node
}

func (b buildData) UpdateState(s any) {
	b.el.State = s
	b.engine.UpdateElement(b.el)
}
