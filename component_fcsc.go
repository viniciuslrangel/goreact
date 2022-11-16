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

func (f *fcscComponent) Keyed(key key, child Node) Node {
	el := f.New(child).(NodeData)
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

func (f *fcscComponent) build(props any) Node {
	return f.target(props.(ChildrenProps).Children[0])
}
