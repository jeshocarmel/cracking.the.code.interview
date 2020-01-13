package ch16

func findContinuousSequence(arr []int) int {

	maxsum := 0
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if maxsum < sum {
			maxsum = sum
		} else if sum < 0 {
			sum = 0
		}
	}

	return maxsum
}
