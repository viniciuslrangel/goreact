package goreact

type RenderInterface interface {
	SetEngine(e *Engine)

	InsertNode(node *NodeData)
	RemoveNode(node *NodeData)

	UpdateElement(node *NodeData)
}

type Engine struct {
	render RenderInterface

	root *NodeData
}

// NewEngine creates a new render engine
//
//goland:noinspection GoExportedFuncWithUnexportedType
func NewEngine(r RenderInterface) *Engine {
	e := &Engine{
		render: r,
	}
	r.SetEngine(e)
	return e
}

func (e *Engine) ParseFragment(node Node) []Node {
	data := node.(*NodeData)
	if data.NativeTyp == "fragment" {
		return data.Children
	}
	return []Node{node}
}

func (e *Engine) DiffChildren(oldChildren []Node, node *NodeData) {
	oldChildrenCount := len(oldChildren)
	childrenCount := len(node.Children)
	if oldChildrenCount == 0 && childrenCount == 0 {
		return
	}
	if oldChildrenCount > childrenCount {
		for i := childrenCount; i < oldChildrenCount; i++ {
			oldChild := oldChildren[i].(*NodeData)
			if oldChild.NativeTyp != "" {
				e.render.RemoveNode(oldChild)
			}
		}
	}
	for i, child := range node.Children {
		if i < oldChildrenCount {
			e.Diff(child.(*NodeData), oldChildren[i].(*NodeData))
		} else {
			newChild := child.(*NodeData)
			newChild.Parent = node
			if newChild.NativeTyp != "" {
				e.render.InsertNode(newChild)
			}
			e.UpdateElement(newChild)
		}
	}
}

func (e *Engine) UpdateElement(el *NodeData) {
	oldChildren := el.Children
	if el.NativeTyp != "" { // Native element
		e.render.UpdateElement(el)
	} else {
		result := el.Typ.build(el.Props).(*NodeData)
		el.Children = e.ParseFragment(result)
	}
	e.DiffChildren(oldChildren, el)
}

func (e *Engine) Diff(element *NodeData, target *NodeData) {
	if !element.sameComp(target) {
		parent := target.Parent

		if target.NativeTyp != "" {
			e.render.RemoveNode(target)
		}

		target.Key = element.Key

		target.NativeTyp = element.NativeTyp
		target.Typ = element.Typ
		target.Props = element.Props

		target.Parent = parent
		target.Children = nil

		target.NativeElement = nil
		target.State = nil

		if target.NativeTyp != "" {
			e.render.InsertNode(target)
		}
		e.UpdateElement(target)
	} else if !element.sameProps(target.Props) {
		target.Props = element.Props
		e.UpdateElement(target)
	}
}

func (e *Engine) Render(element Node) {
	e.root = new(NodeData)
	e.Diff(element.(*NodeData), e.root)
}
