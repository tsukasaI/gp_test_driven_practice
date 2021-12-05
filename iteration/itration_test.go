package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 8)
	fmt.Println(result)
	// Output: bbbbbbbb
}

func TestContain(t *testing.T) {
	a := "foo"
	b := "seafood"
	result := strings.Contains(b, a)

	if !result {
		t.Errorf("%q does not contain %q", b, a)
	}
}
