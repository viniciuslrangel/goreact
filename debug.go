package goreact

import (
	"fmt"
	"strings"
)

func DumpTree(root Node) string {
	var sb strings.Builder
	var draw func(el *NodeData, padding int)
	draw = func(el *NodeData, padding int) {
		spacer := "└─"
		if padding > 0 {
			spacer = strings.Repeat("  ", padding) + spacer
		}
		var data any
		var name string
		var key string
		var children []Node

		if el.Key.Has {
			key = fmt.Sprintf(" key=%v", el.Key.Key)
		}

		if cp, ok := el.Props.(childrenProps); ok {
			children = cp.GetChildren()
		}
		comp := el.Typ
		if comp == nil {
			if el.NativeTyp != "" {
				name = el.NativeTyp
				data = el.Props
			} else {
				name = "nil"
				data = "nil"
			}
		} else {
			name = comp.getName()
			data = el.Props
			var build buildData
			build.el = el
			value := comp.build(build)
			if len(children) == 0 {
				children = append(children, value)
			}
		}
		sb.WriteString(fmt.Sprintf("%s%s %+v%s", spacer, name, data, key))
		for _, child := range children {
			sb.WriteString("\n")
			draw(child.(*NodeData), padding+1)
		}
	}
	draw(root.(*NodeData), 0)
	return sb.String()
}
