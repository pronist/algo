package deepEqual

import (
	"reflect"
	"testing"
)

type mapTest struct {
	a, b map[string]int
	eq bool
}

var mapTests = []mapTest{
	{map[string]int{"a": 1}, map[string]int{"b": 1}, true},
	{map[string]int{"a": 1}, map[string]int{"a": 1}, false},
}

func TestMapTable(t *testing.T) {
	for _ ,test := range mapTests {
		if r := reflect.DeepEqual(test.a, test.b); r != test.eq {
			t.Errorf("when a = %#v and b = %#v, want %t, got %t", test.a, test.b, r, test.eq)
		}
	}
}