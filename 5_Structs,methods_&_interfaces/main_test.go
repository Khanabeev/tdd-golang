package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g",shape, got, want)
		}
	}

	areaTests := []struct {
		name string
		shape Shape
		want  float64
	}{
		{name: "Area Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Area Circle", shape: Circle{radios: 10}, want: 314.1592653589793},
		{name: "Area Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
