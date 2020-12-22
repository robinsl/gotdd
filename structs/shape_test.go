package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("get rectangles area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("get circle area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})


	t.Run("tableDriven test", func(t *testing.T) {
		areaTests := []struct{
			name string
			shape Shape
			want float64
		}{
			{"Rectangle", Rectangle{12, 6}, 72.0},
			{"Circle", Circle{10}, 314.1592653589793},
			{"Triangle", Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()
				if got != tt.want {
					t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
				}
			})
		}
	})
}
