package goreact

// FCS Functional Component with State without Props
//
//goland:noinspection GoExportedFuncWithUnexportedType
func FCS[S any](name string, defaultState S, target func(state S, updateState func(S)) Node) *fcsComponentNoProps[S] {
	return &fcsComponentNoProps[S]{
		name:         name,
		defaultState: defaultState,
		target:       target,
	}
}

type fcsComponentNoProps[S any] struct {
	name         string
	defaultState S
	target       func(state S, updateState func(S)) Node
}

func (f *fcsComponentNoProps[S]) Keyed(key Key) Node {
	el := f.New().(*NodeData)
	el.Key = key
	return el
}

func (f *fcsComponentNoProps[S]) New() Node {
	return &NodeData{
		Typ:   f,
		State: f.defaultState,
	}
}

func (f *fcsComponentNoProps[S]) getName() string {
	return f.name
}

func (f *fcsComponentNoProps[S]) build(data buildData) Node {
	return f.target(data.el.State.(S), func(newState S) {
		data.el.State = newState
		data.engine.UpdateElement(data.el)
	})
}
