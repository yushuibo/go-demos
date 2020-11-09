/*
 * @Since: 2020-02-23 00:08:18
 * @Author: shy
 * @Email: yushuibo@ebupt.com / hengchen2005@gmail.com
 * @Version: v1.0
 * @Description: -
 */
package main

import "fmt"

func MinAndMax(nums ...int) (min, max int) {
	if len(nums) < 2 {
		fmt.Printf("This function except at least 2 arguments, but %d gaven.", len(nums))
	}

	max = nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	min = nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return
}

func main() {
	nums := []int{5, 32, 4, 22, 1}
	// MinAndMax(1)
	min, max := MinAndMax(nums...)
	fmt.Printf("Min=%d, Max=%d", min, max)
}
