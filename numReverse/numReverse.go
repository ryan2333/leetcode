package main

func NumReverse(n int) int {
	max := 1<<31 - 1
	min := max * -1

	//fmt.Println(math.MaxUint32, math.MinInt32)
	//fmt.Println(max, min)
	var rn int
	for n != 0 {
		rn = rn*10 + n%10
		if rn > max || rn < min+1 {
			return 0
		}
		n = n / 10
	}

	return rn
}
