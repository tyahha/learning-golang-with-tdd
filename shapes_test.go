package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(&Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"circle", Circle{10.0}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v, got %g want %g", test.shape, got, test.hasArea)
			}
		})
	}
}
