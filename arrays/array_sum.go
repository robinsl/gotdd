package arrays

func ArraySum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}
