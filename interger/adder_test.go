package interger

import "testing"

func TestAdder(t *testing.T)  {
	sum :=Add(2,2)
	expected :=4
	if sum != expected{
		t.Errorf("sum '%d' ,expected '%d'",sum,expected)
	}
}

func Add(x,y int) int {
	return x+y
}