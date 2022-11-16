package main

import (
	"fmt"
	. "github.com/viniciuslrangel/goreact"
)

type ButtonProps struct {
	Label string
	Flat  bool
}

var Button = FCP("Button", func(props ButtonProps) Node {
	return NativeEl("button", map[string]any{
		"label": props.Label,
		"flat'": props.Flat,
	})
})

type TextProps struct {
	Label string
}

var Text = FCP("Text", func(props TextProps) Node {
	return NativeEl("text", map[string]any{
		"label": props.Label,
	})
})

var Column = FCC("Column", func(children ...Node) Node {
	return NativeEl("column", ChildrenProps{
		Children: children,
	})
})

var HelloWorld = FC("HelloWorld", func() Node {
	return Column.New(
		Text.Keyed(
			Key(1),
			TextProps{
				Label: "Hello World",
			},
		),
		Button.New(ButtonProps{
			Label: "Click me",
		}),
	)
})

var App = FC("App", func() Node {
	return HelloWorld.New()
})

func main() {
	a := DumpTree(App.New())
	//a := App()
	fmt.Println(a)
}
