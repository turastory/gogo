package basic

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 6)
	expected := "aaaaaa"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
