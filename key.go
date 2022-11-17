package goreact

type Key struct {
	Key uint64
	Has bool
}

var NoKey = Key{Has: false}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewKey(k uint64) Key {
	return Key{Key: k, Has: true}
}
