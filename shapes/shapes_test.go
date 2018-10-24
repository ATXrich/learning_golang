package shapes

import "testing"

func TestArea(t *testing.T) {
	want := 100.0
	got := Area(10.0, 10.0)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
