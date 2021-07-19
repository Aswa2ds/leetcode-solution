package offer

//func hammingWeight(num uint32) int {
//	var hamming uint32
//	for num > 0 {
//		hamming += num & 1
//		num = num >> 1
//	}
//	return int(hamming)
//}

func hammingWeight(num uint32) int {
	var hamming int
	for ; num > 0; num = num & (num-1) {
		hamming++
	}
	return hamming
}
