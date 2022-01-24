package shapes

import "testing"

func assertEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assertEquals(t, got, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Circle{10}, 314.1592653589793},
		{Rectangle{12, 6}, 72.0},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertEquals(t, got, tt.want)
	}
}
