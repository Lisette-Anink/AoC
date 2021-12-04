package utils

func TransposeInt(array [][]int) map[int][]int {
	transposed := map[int][]int{}
	for i := 0; i < len(array[0]); i++ {
		transposed[i] = []int{}
		for _, l := range array {
			transposed[i] = append(transposed[i], l[i])
		}
	}
	return transposed
}
func TransposeString(array [][]string) map[int][]string {
	transposed := map[int][]string{}
	for i := 0; i < len(array[0]); i++ {
		transposed[i] = []string{}
		for _, l := range array {
			transposed[i] = append(transposed[i], l[i])
		}
	}
	return transposed
}
