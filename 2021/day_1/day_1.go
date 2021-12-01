package day_one

func CountIncreased(a []int) int {
	count := 0
	for i, n := range a {
		if i != 0 && n > a[i-1] {
			count++
		}
	}
	return count
}

func CountIncreasedSlidingWindow(a []int) int {
	count := 0
	for i := range a {
		if i >= 3 && (a[i-3]+a[i-2]+a[i-1]) < (a[i-2]+a[i-1]+a[i]) {
			count++
		}
	}
	return count
}
