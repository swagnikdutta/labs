package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := rectangle.Perimeter()
	want := 60.0

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %g, got %g", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		checkArea(t, rectangle, 200.0)
	})

	t.Run("cicles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// Table driven tests are useful when you want to build a list of
// test cases that can be tested in the same manner.
//
// Table driven tests are very useful when you want to test various
// implementations of an interface, or if the data being passed in to
// a function has lots of different requirements that need to be
// tested. It's kind of a more thorough testing tool.
func TestAreaV2(t *testing.T) {
	// creating an anonymous struct
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 200.0},
		{Circle{10}, 314.1592653589793},
		// Now it's very easy for a developer to add a test case for a different shape
		{shape: Triangle{12, 6}, want: 36},
	}

	// iterate over the cases
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v want %g, got %g", tt.shape, tt.want, got)
		}
	}
}

func TestAreaV3(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, hasArea: 200.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36},
	}

	// iterate over the cases
	for _, tt := range areaTests {
		// Using tt.name from the case to use it as the `t.Run` test name.
		// Wrapping each test inside t.Run gives clearer test output on failures
		// as it will print the name of the case.
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v want %g, got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
