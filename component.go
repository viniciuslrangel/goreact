package goreact

type Component interface {
	getName() string

	build(props any) Node
}
