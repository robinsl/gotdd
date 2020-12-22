package arrays

func ArraySumAll(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		sums = append(sums, ArraySum(number))
	}
	return
}
