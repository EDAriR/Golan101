package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

/*
数组有一个有趣的属性，它的大小也属于类型的一部分，如果你尝试将 [4]int 作为 [5]int 类型的参数传入函数，是不能通过编译的。它们是不同的类型，就像尝试将 string 当做 int 类型的参数传入函数一样。
因为这个原因，所以数组比较笨重，大多数情况下我们都不会使用它。
Go 的切片（slice）类型不会将集合的长度保存在类型中，因此它的尺寸可以是不固定的。

*/
