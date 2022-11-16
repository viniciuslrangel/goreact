package goreact

// FC Functional Component without Props
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FC(name string, target func() Node) *fcComponentNoProps {
	return &fcComponentNoProps{
		name:   name,
		target: target,
	}
}

type fcComponentNoProps struct {
	name   string
	target func() Node
}

func (f *fcComponentNoProps) Keyed(key key) Node {
	el := f.New().(*NodeData)
	el.Key = key
	return el
}

func (f *fcComponentNoProps) New() Node {
	return &NodeData{
		Typ: f,
	}
}

func (f *fcComponentNoProps) getName() string {
	return f.name
}

func (f *fcComponentNoProps) build(props any) Node {
	return f.target()
}
