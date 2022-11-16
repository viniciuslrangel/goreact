package goreact

// FCSP Functional Component with State and Props
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FCSP[S, T any](name string, defaultState S, target func(props T, state S, updateState func(S)) Node) *fcsComponent[S, T] {
	return &fcsComponent[S, T]{
		name:         name,
		defaultState: defaultState,
		target:       target,
	}
}

type fcsComponent[S, T any] struct {
	name         string
	defaultState S
	target       func(props T, state S, updateState func(S)) Node
}

func (f *fcsComponent[S, T]) Keyed(key key, props T) Node {
	el := f.New(props).(*NodeData)
	el.Key = key
	return el
}

func (f *fcsComponent[S, T]) New(props T) Node {
	return &NodeData{
		Typ:   f,
		Props: props,
		State: f.defaultState,
	}
}

func (f *fcsComponent[S, T]) getName() string {
	return f.name
}

func (f *fcsComponent[S, T]) build(data buildData) Node {
	return f.target(data.el.Props.(T), data.el.State.(S), func(newState S) {
		data.el.State = newState
		data.engine.UpdateElement(data.el)
	})
}
