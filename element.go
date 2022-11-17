package goreact

type Node interface {
	_nodeData()
}

type NodeData struct {
	IsDirty bool
	Key     Key

	NativeTyp string
	Typ       Component
	Props     any

	Parent   *NodeData
	Children []Node

	NativeElement any

	State any
}

func (e *NodeData) _nodeData() {}

func (e *NodeData) sameComp(el *NodeData) bool {
	if e.Key != el.Key {
		return false
	}
	if e.NativeTyp != "" && e.NativeTyp != el.NativeTyp {
		return false
	}
	if e.Typ != el.Typ {
		return false
	}
	return true
}

func (e *NodeData) sameProps(props any) bool {
	return compareProps(e.Props, props)
}

func NativeEl(typ string, props any) Node {
	return &NodeData{
		NativeTyp: typ,
		Props:     props,
	}
}

func Fragment(children ...Node) Node {
	return &NodeData{
		NativeTyp: "fragment",
		Props: ChildrenProps{
			Children: children,
		},
	}
}
