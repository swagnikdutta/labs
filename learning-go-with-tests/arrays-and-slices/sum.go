package arrays_and_slices

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			tail := slice[1:]
			sum := Sum(tail)
			result = append(result, sum)
		}
	}
	return result
}
