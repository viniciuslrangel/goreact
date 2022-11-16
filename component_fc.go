package goreact

// FCP Functional Component with Props
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FCP[T any](name string, target func(props T) Node) *fcComponent[T] {
	return &fcComponent[T]{
		name:   name,
		target: target,
	}
}

type fcComponent[T any] struct {
	name   string
	target func(props T) Node
}

func (f *fcComponent[T]) Keyed(key key, props T) Node {
	el := f.New(props).(*NodeData)
	el.Key = key
	return el
}

func (f *fcComponent[T]) New(props T) Node {
	return &NodeData{
		Typ:   f,
		Props: props,
	}
}

func (f *fcComponent[T]) getName() string {
	return f.name
}

func (f *fcComponent[T]) build(props any) Node {
	return f.target(props.(T))
}
