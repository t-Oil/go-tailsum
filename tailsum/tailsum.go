package tailsum

func TailSum(input []int, sum int) int {
	testInput := input
	k := sum

	reverseInput := reverseSlice(testInput)

	res := isTailSum(reverseInput, k)

	return res
}

func reverseSlice(input []int) []int {
	var res []int
	for i := range input {
		res = append(res, input[len(input)-1-i])
	}

	return res
}

func isTailSum(input []int, k int) int {
	var sum int
	var lenIndex int

	if len(input) == 0 {
		return 0
	}

	for _, val := range input {
		sum += val
		lenIndex++

		if sum == k {
			return lenIndex
		}
	}

	return isTailSum(input[:len(input)-1], k)
}
