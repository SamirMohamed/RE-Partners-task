package store

import (
	"reflect"
	"testing"
)

func TestGetNumOfPacksSizes(t *testing.T) {
	testCases := []struct {
		packsSizes              []int
		orderItemsCount         int
		expectedNumOfPacksSizes map[int]int
	}{
		{
			packsSizes:              []int{1000, 500, 2000, 250, 5000},
			orderItemsCount:         1,
			expectedNumOfPacksSizes: map[int]int{250: 1},
		},
		{
			packsSizes:              []int{250, 500, 1000, 2000, 5000},
			orderItemsCount:         250,
			expectedNumOfPacksSizes: map[int]int{250: 1},
		},
		{
			packsSizes:              []int{250, 500, 1000, 2000, 5000},
			orderItemsCount:         251,
			expectedNumOfPacksSizes: map[int]int{500: 1},
		},
		{
			packsSizes:              []int{250, 500, 1000, 2000, 5000},
			orderItemsCount:         501,
			expectedNumOfPacksSizes: map[int]int{250: 1, 500: 1},
		},
		{
			packsSizes:              []int{250, 500, 1000, 2000, 5000},
			orderItemsCount:         12001,
			expectedNumOfPacksSizes: map[int]int{250: 1, 5000: 2, 2000: 1},
		},
		{
			packsSizes:              []int{250, 500, 1000, 2000, 5000},
			orderItemsCount:         4749,
			expectedNumOfPacksSizes: map[int]int{2000: 2, 500: 1, 250: 1},
		},
		{
			packsSizes:              []int{250},
			orderItemsCount:         251,
			expectedNumOfPacksSizes: map[int]int{250: 2},
		},
		{
			packsSizes:              []int{250},
			orderItemsCount:         0,
			expectedNumOfPacksSizes: map[int]int{},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := getNumOfPacksSizes(tc.packsSizes, tc.orderItemsCount)
			if !reflect.DeepEqual(tc.expectedNumOfPacksSizes, got) {
				t.Fatalf("expected: %v, got: %v", tc.expectedNumOfPacksSizes, got)
			}
		})
	}
}
