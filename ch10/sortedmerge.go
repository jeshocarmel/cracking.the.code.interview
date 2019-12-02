package ch10

func merge(arr1, arr2 []int, lastA, lastB int) {

	indexA := lastA - 1
	indexB := lastB - 1

	indexMerged := lastB + lastA - 1

	// fmt.Println(indexA, indexB, indexMerged)

	for indexB >= 0 {

		if indexA >= 0 && arr1[indexA] > arr2[indexB] {
			arr1[indexMerged] = arr1[indexA]
			indexA--
		} else {
			arr1[indexMerged] = arr2[indexB]
			indexB--
		}
		indexMerged--
	}
}
