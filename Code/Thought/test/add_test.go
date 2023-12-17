package test

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{3000, 4000, 5000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d); "+
				"expected %d, got %d", tt.a, tt.b, tt.c, actual)
		}
	}
}
func BenchmarkTriangle(bb *testing.B) {
	a, b, c := 3000, 4000, 5000
	for i := 0; b < bb.N; i++ {
		if actual := calcTriangle(a, b); actual != c {
			bb.Errorf("calcTriangle(%d,%d); "+
				"expected %d, got %d", a, b, c, actual)
		}
	}

}
