package goreact

var NoKey = key{Has: false}

//goland:noinspection GoExportedFuncWithUnexportedType
func Key(k uint64) key {
	return key{Key: k, Has: true}
}
