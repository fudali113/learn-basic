package main

func sort(array []int) {
	arrayLen := len(array)
	for gap := arrayLen / 2; gap > 0; gap /= 2 {
		for i := gap; i < arrayLen; i++ {
			iv := array[i]
			for j := i; j-gap > 0 && array[j-gap] > iv; j-- {
				array[j-gap], array[j] = iv, array[j-gap]
			}
		}
	}
}
func main() {
	array := []int{10, 14, 73, 25, 23, 13, 27, 94, 33, 39, 25, 59, 94, 65, 82, 45}
	sort(array)
	for _, v := range array {
		println(v)
	}
}
