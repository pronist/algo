package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	r := Sum(1, 2)
	if r != 3 {
		t.Fatal("sum(1, 2) should be 3, but doesn't match")
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2))
	// Output: 3
}
