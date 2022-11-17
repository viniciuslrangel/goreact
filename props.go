package goreact

import (
	"github.com/modern-go/reflect2"
	"reflect"
)

type IChildrenProps interface {
	GetChildren() []Node
}

type ChildrenProps struct {
	Children []Node
}

func (p ChildrenProps) GetChildren() []Node {
	return p.Children
}

func (p Key) GetKey() (uint64, bool) {
	return p.Key, p.Has
}

var compareFuncCache = make(map[reflect2.Type]func(a, b any) bool)

func genCompareFunc(t reflect2.Type) (output func(a, b any) bool) {
	defer func() {
		if output != nil {
			compareFuncCache[t] = output
		}
	}()
	if t.Kind() == reflect.Struct {
		t := t.(reflect2.StructType)
		var fieldList []reflect2.StructField
		numField := t.NumField()
		for i := 0; i < numField; i++ {
			field := t.Field(i)
			if field.Name() == "Children" {
				continue
			}
			fieldList = append(fieldList, field)
		}
		return func(a, b any) bool {
			pa := reflect2.PtrOf(a)
			pb := reflect2.PtrOf(b)
			for _, field := range fieldList {
				// TODO Nested struct?
				/*if !compareProps(a.FieldByName(field.Name).Interface(), b.FieldByName(field.Name).Interface()) {
					return false
				}*/
				fieldType := field.Type()
				fa := fieldType.PackEFace(field.UnsafeGet(pa))
				fb := fieldType.PackEFace(field.UnsafeGet(pb))
				if fa != fb {
					return false
				}
			}
			return true
		}
	} else {
		// TODO
		panic("not supported. Type: " + t.Kind().String())
		return nil
	}
}

func compareProps(x any, y any) bool {
	if x == nil || y == nil {
		return x == y
	}

	typ1 := reflect2.TypeOf(x)
	if typ1 != reflect2.TypeOf(y) {
		return false
	}
	compare, ok := compareFuncCache[typ1]
	if !ok {
		compare = genCompareFunc(typ1)
		compareFuncCache[typ1] = compare
	}
	return compare(x, y)
}
