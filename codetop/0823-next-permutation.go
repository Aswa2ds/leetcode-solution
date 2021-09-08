package codetop

import "fmt"

func nextPermutation(nums []int)  {
	n := len(nums)
	var index int = -1
	for i := n-1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i-1
			break
		}
	}
	fmt.Println(index)
	if index >= 0 {
		for i := n-1; i > index; i-- {
			if nums[i] > nums[index] {
				nums[i], nums[index] = nums[index], nums[i]
				break
			}
		}
	}
	fmt.Println(nums)
	func(list []int ){
		for i := 0; i < len(list)/2; i++ {
			list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
		}
	}(nums[index+1:])
}
