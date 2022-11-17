package goreact

// FCC Functional Component with Children
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FCC(name string, target func(children ...Node) Node) *fccComponent {
	return &fccComponent{
		name:   name,
		target: target,
	}
}

type fccComponent struct {
	name   string
	target func(children ...Node) Node
}

func (f *fccComponent) Keyed(key Key, children ...Node) Node {
	el := f.New(children...).(*NodeData)
	el.Key = key
	return el
}

func (f *fccComponent) New(children ...Node) Node {
	var props ChildrenProps
	props.Children = children
	return &NodeData{
		Typ:   f,
		Props: props,
	}
}

func (f *fccComponent) getName() string {
	return f.name
}

func (f *fccComponent) build(data buildData) Node {
	return f.target(data.el.Props.(ChildrenProps).Children...)
}
