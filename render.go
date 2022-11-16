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
		childData := child.(*NodeData)
		if i < oldChildrenCount {
			oldChildData := oldChildren[i].(*NodeData)
			e.Diff(childData, oldChildData)
			*childData = *oldChildData
		} else {
			newChild := childData
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
		var data buildData
		data.engine = e
		data.el = el
		result := el.Typ.build(data).(*NodeData)
		result.IsDirty = false
		el.Children = e.ParseFragment(result)
		for _, child := range el.Children {
			childData := child.(*NodeData)
			childData.IsDirty = true
		}
	}
	e.DiffChildren(oldChildren, el)
}

func (e *Engine) Diff(NewNode *NodeData, currentNode *NodeData) {
	shouldUpdate := false
	isNew := false
	if !NewNode.sameComp(currentNode) {
		shouldUpdate = true
		isNew = true
	} else if NewNode.IsDirty || !NewNode.sameProps(currentNode.Props) {
		shouldUpdate = true
	}

	if shouldUpdate {
		parent := currentNode.Parent

		if isNew {
			if currentNode.NativeTyp != "" {
				e.render.RemoveNode(currentNode)
			}
			currentNode.NativeTyp = NewNode.NativeTyp
			currentNode.Typ = NewNode.Typ
			currentNode.Parent = parent
			currentNode.Children = nil
		} else if currentNode.NativeElement != nil {

		}

		currentNode.Key = NewNode.Key
		currentNode.Props = NewNode.Props
		currentNode.State = NewNode.State

		if isNew {
			if currentNode.NativeTyp != "" {
				e.render.InsertNode(currentNode)
			}
		}
		e.UpdateElement(currentNode)
	}
}

func (e *Engine) Render(element Node) {
	e.root = new(NodeData)
	e.Diff(element.(*NodeData), e.root)
}
