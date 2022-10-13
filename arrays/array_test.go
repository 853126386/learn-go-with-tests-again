package arrays

import "testing"

func TestSum(t *testing.T)  {
	//var numbers [5]int=[5]int{1,2,3,4,5}
	//var numbers =[5]int{1,2,3,4,5}
	//numbers:=[5]int{1,2,3,4,5} //确定长度的数组
	numbers:=[...]int{1,2,3,4,5}  //不确定长度的数组
	sum :=Sum(numbers)
	expected :=15
	if sum != expected{
		t.Errorf("sum '%d' but expected '%d' ,given '%v'",sum,expected,numbers)
	}
}
