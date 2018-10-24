package shapes

import "testing"

func TestArea(t *testing.T) {

	testTable := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range testTable {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}

}
