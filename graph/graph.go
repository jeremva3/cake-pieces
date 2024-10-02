package graph

func RemoveZeroSumConsecutive(numbers []int) []int {
	if numbers == nil {
		return []int{}
	}

	prefixSumIndexMap := make(map[int]int)
	result := []int{}
	currentSum := 0

	prefixSumIndexMap[0] = -1

	for index := 0; index < len(numbers); index++ {
		currentSum += numbers[index]

		if startIndex, found := prefixSumIndexMap[currentSum]; found {
			result = result[:startIndex+1]
		} else {
			prefixSumIndexMap[currentSum] = len(result)
			result = append(result, numbers[index])
		}
	}

	return result
}
