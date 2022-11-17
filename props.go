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

	type DeepStructField struct {
		reflect2.StructField
		compareFunc func(a, b any) bool
	}

	if t.Kind() == reflect.Struct {
		t := t.(reflect2.StructType)

		var fieldList []reflect2.StructField
		var fieldStructList []DeepStructField

		numField := t.NumField()
		for i := 0; i < numField; i++ {
			field := t.Field(i)
			if field.Name() == "Children" {
				continue
			}

			if field.Type().Kind() == reflect.Struct {
				fieldStructList = append(fieldStructList, DeepStructField{
					StructField: field,
					compareFunc: genCompareFunc(field.Type()),
				})
			} else {
				fieldList = append(fieldList, field)
			}
		}
		return func(a, b any) bool {
			pa := reflect2.PtrOf(a)
			pb := reflect2.PtrOf(b)
			for _, field := range fieldList {
				fieldType := field.Type()
				fa := fieldType.PackEFace(field.UnsafeGet(pa))
				fb := fieldType.PackEFace(field.UnsafeGet(pb))
				if fa != fb {
					return false
				}
			}
			for _, field := range fieldStructList {
				fieldType := field.Type()
				fa := fieldType.PackEFace(field.UnsafeGet(pa))
				fb := fieldType.PackEFace(field.UnsafeGet(pb))
				if !field.compareFunc(fa, fb) {
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
	}
	return compare(x, y)
}
