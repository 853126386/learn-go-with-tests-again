package _struct

import "testing"



func TestPerimeter(t *testing.T){
	t.Run("Rectangle of Perimeter", func(t *testing.T) {
		rectangle :=Rectangle{10.0,10.0}
		got :=Perimeter(rectangle)
		want :=40.0
		if got!=want{
			t.Errorf("got %.2f  want %.2f",got,want)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T,shape Shape,want float64) {
		t.Helper()
		got :=shape.Area()
		if got != want{
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Rectangle of Perimeter", func(t *testing.T) {
		rectangle :=Rectangle{10.0,10.0}
		want :=100.0
		checkArea(t,rectangle,want)
	})


	t.Run("circle of Perimeter", func(t *testing.T) {
		circle :=Circle{10.0}
		want :=40.0
		checkArea(t,circle,want)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{shape:Rectangle{Width: 10.0,Height: 10.0},want: 100.0},
		{shape: Circle{Radius: 10.0},want:314.1592653589793},
		{shape: Triangle{BottomLength: 10.0,Height: 2.0},want: 10.0},
	}
	for _,tt :=range areaTests{
		got :=tt.shape.Area()
		if got != tt.want{
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}


}
