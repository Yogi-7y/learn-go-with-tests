package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	result := make([]int, len(numbers))

	for index, numberArray := range numbers {
		result[index] = Sum(numberArray)
	}

	return result
}
