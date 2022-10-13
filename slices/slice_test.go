package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {
	checkSum := func(t *testing.T,got,want []int) {
		t.Helper()
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v but want %v",got,want)
		}
	}
	t.Run("collection of any size", func(t *testing.T) {
		numbers:=[]int{1,2,3,4,5}  //不确定长度的数组
		sum :=Sum(numbers)
		expected :=15
		if sum != expected{
			t.Errorf("sum '%d' but expected '%d' ,given '%v'",sum,expected,numbers)
		}
	})

	t.Run("sum of all", func(t *testing.T) {
		sumAll :=SumAll([]int{1,2}, []int{0,9})
		expected :=[]int{3, 9}
		checkSum(t,sumAll,expected)

	})



	t.Run("sum of all tails", func(t *testing.T) {
		sumAll :=SumAll([]int{1,2}, []int{0,9})
		expected :=[]int{3, 9}
		checkSum(t,sumAll,expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got :=SumAllTails([]int{},[]int{1,3,5})
		want :=[]int{0,9}
		checkSum(t,got,want)
	})

}

