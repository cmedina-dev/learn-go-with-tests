package structs_methods_interfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func BenchmarkRectangle_Area(b *testing.B) {
	rectangle := Rectangle{10, 10}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{Width: 10, Height: 10}
	fmt.Println(rectangle.Area())
	// Output: 100
}

func BenchmarkCircle_Area(b *testing.B) {
	circle := Circle{10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}

func ExampleCircle_Area() {
	circle := Circle{Radius: 10}
	fmt.Println(circle.Area())
	// Output: 314.1592653589793
}

func BenchmarkTriangle_Area(b *testing.B) {
	triangle := Triangle{Base: 10, Height: 10}
	for i := 0; i < b.N; i++ {
		triangle.Area()
	}
}

func ExampleTriangle_Area() {
	triangle := Triangle{Base: 10, Height: 10}
	fmt.Println(triangle.Area())
	// Output: 50
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 3, Height: 2}, want: 3.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			assertCorrectGeometry(t, got, want, tt.shape)
		})
	}
}

func assertCorrectGeometry(t testing.TB, got, want float64, shape Shape) {
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}
