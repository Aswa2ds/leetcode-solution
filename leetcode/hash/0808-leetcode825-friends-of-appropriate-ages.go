package hash

func numFriendRequests(ages []int) int {
	var ageRange = 120

	ageCount := make([]int, ageRange+1)
	sum := make([]int, ageRange+1)
	for _, age := range ages {
		ageCount[age]++
	}

	for i := 1; i <= ageRange; i++ {
		sum[i] = sum[i-1] + ageCount[i]
	}

	res := 0
	for age, count := range ageCount {
		if age <= 14 {
			continue
		}
		minAge := age>>1 + 7 // target age should be bigger than minAge
		maxAge := age // target age should be less or equal to maxAge
		res += count * (sum[maxAge] - sum[minAge]-1)
	}
	return res
}
