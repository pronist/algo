package bench

import (
	"reflect"
	"testing"
)

type StructAccess struct {
	Int int
}

func BenchmarkDetectTypeReflect(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		rv := reflect.ValueOf(s)
		if rv.Type().Name() == "StructAccess" {
			_ = s.(StructAccess).Int
		}
	}
}

func BenchmarkDetectTypeNone(b *testing.B) {
	s := StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		_ = s.Int
	}
}
