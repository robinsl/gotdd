package arrays

func ArraySumTrail(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		sums = append(sums, ArraySum(number[1:]))
	}
	return
}
