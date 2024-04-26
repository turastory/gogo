package basic

import "testing"

func assertFloat(got float64, want float64, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{2.0, 5.0})
	want := 14.0
	assertFloat(got, want, t)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{width: 2.0, height: 5.0}, 10.0},
		{"Circle", Circle{radius: 10.0}, 314.1592653589793},
		{"Triangle", Triangle{base: 3.0, height: 4.0}, 6.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			area := test.shape.Area()
			assertFloat(area, test.want, t)
		})
	}
}
