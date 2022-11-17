package goreact

// FCSC Functional Component with a single child
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FCSC(name string, target func(child Node) Node) *fcscComponent {
	return &fcscComponent{
		name:   name,
		target: target,
	}
}

type fcscComponent struct {
	name   string
	target func(child Node) Node
}

func (f *fcscComponent) Keyed(key Key, child Node) Node {
	el := f.New(child).(*NodeData)
	el.Key = key
	return el
}

func (f *fcscComponent) New(child Node) Node {
	var props ChildrenProps
	props.Children = []Node{child}
	return &NodeData{
		Typ:   f,
		Props: props,
	}
}

func (f *fcscComponent) getName() string {
	return f.name
}

func (f *fcscComponent) build(data buildData) Node {
	return f.target(data.el.Props.(ChildrenProps).Children[0])
}
