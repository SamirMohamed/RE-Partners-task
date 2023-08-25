package store

import "sort"

func GetNumOfPacks(packsSizes []int, orderItemsCount int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packsSizes)))

	numOfPacksSizes := make(map[int]int)
	lowestPackSize := packsSizes[len(packsSizes)-1]

	for idx := 0; idx < len(packsSizes); idx++ {
		currPackSize := packsSizes[idx]
		if orderItemsCount > (currPackSize - lowestPackSize) {
			if orderItemsCount > currPackSize {
				numOfPacksSizes[currPackSize] = orderItemsCount / currPackSize
				orderItemsCount = orderItemsCount % currPackSize
				idx--
			} else {
				numOfPacksSizes[currPackSize]++
				orderItemsCount = 0
				break
			}
		}
	}

	return numOfPacksSizes
}
