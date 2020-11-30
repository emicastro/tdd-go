package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertRepeated := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("when receives 'a' adn 3 returns 'aaaaa'", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertRepeated(t, expected, repeated)
	})

	t.Run("when receives 'b' and 5 returns 'bbbbb'", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"
		assertRepeated(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
