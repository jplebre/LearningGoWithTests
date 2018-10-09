package structs

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestPerimeter(t *testing.T) {
	g := NewGomegaWithT(t)

	// Let's refactor the 2 tests, by using an interface called Shape
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		g.Expect(got).To(Equal(want))
	}

	t.Run("Calculates perimeter of Rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, 40)
	})

	t.Run("Calculates perimeter of Circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 62.83185307179586)
	})
}

// Let's refactor the previous test to use Tables
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Calculates area of Rectangles", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Calculates area of Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Calculates area of Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v %.2f want %.2f", tt.shape, got, tt.hasArea)
		}
	}
}
