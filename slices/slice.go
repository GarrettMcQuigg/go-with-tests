package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(num ...[]int) []int {
	var sums []int
	for _, n := range num {
		sums = append(sums, Sum(n))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {

	var sum []int
	for _, num := range numbers {
		if len(num) == 0 {
			sum = append(sum, 0)
		} else {
			tail := num[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
