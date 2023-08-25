package store

import (
	"context"
	"errors"
	"gymshark/cache"
	"sort"
	"strconv"
)

const PacksSizesCacheKey = "packs_sizes"

func SeedData() {
	cache.CacheClient.SAdd(context.TODO(), PacksSizesCacheKey, 250, 500, 1000, 2000, 5000)
}

func GetNumOfPacks(orderItemsCount int) (map[int]int, error) {
	packsSizes, err := getPacksSizes()
	if err != nil {
		return nil, err
	}

	if len(packsSizes) < 1 {
		return nil, errors.New("empty packs")
	}

	return getNumOfPacksSizes(packsSizes, orderItemsCount), nil
}

func getNumOfPacksSizes(packsSizes []int, orderItemsCount int) map[int]int {
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

func getPacksSizes() ([]int, error) {
	// For seek of simplicity, I used redis to store/get the packs.
	// We may use dbs for permanent data.

	cmd := cache.CacheClient.SMembers(context.TODO(), PacksSizesCacheKey)
	packsSizesStr, err := cmd.Result()
	if err != nil {
		return []int{}, nil
	}

	var packsSizesInt []int

	for _, pst := range packsSizesStr {
		psi, err := strconv.Atoi(pst)
		if err != nil {
			return []int{}, err
		}
		packsSizesInt = append(packsSizesInt, psi)
	}

	return packsSizesInt, nil
}
