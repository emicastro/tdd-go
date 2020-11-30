package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertSum := func(t *testing.T, expected, sum int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("Adds 2 + 2, expects 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertSum(t, expected, sum)
	})

	t.Run("Adds 5 + 8, expects 13", func(t *testing.T) {
		sum := Add(5, 8)
		expected := 13
		assertSum(t, expected, sum)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
