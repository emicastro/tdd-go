package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assertEquals := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		got := rectangle.Perimeter()
		want := 40.0

		assertEquals(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Perimeter()
		want := 62.831853071795865

		assertEquals(t, got, want)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v has area %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
