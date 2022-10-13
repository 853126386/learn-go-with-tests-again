package slices

func Sum(numbers []int)(sum int)  {
	//for i:=0;i< len(numbers);i++ {
	//	sum += numbers[i]
	//}

	//range 会迭代数组，每次迭代都会返回数组元素的索引和值。我们选择使用 _ 空白标志符 来忽略索引。
	for _,number := range numbers{
		sum += number
	}
	return

}

func SumAll(numbersToSum ...[]int) []int{
	sumAll := make([]int,len(numbersToSum),len(numbersToSum))
	for index,numbers :=range numbersToSum{
		sum :=0
		for _,value :=range numbers{
			sum+=value
		}
		sumAll[index]=sum
	}
	return sumAll
}


func SumAllTails(numbersToSum ...[]int) (sumAll[]int){
	for _,numbers :=range numbersToSum{
		tail :=[]int{0}
		if len(numbers) > 0{
			tail =numbers[1:]
		}
		sumAll=append(sumAll,Sum(tail))
	}
	return
}
