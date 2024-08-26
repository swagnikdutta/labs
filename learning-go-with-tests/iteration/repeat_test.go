package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 10 times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("repeat 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("d", 5)
	fmt.Println(repeated)
	// Output: ddddd
}
