package shell

func sort(array []int) {
	n := 0
	arrayLen := len(array)
	for gap := arrayLen / 2; gap > 0; gap /= 2 {
		for i := gap; i < arrayLen; i++ {
			iv := array[i]
			for j := i; j-gap > 0 && array[j-gap] > iv; j -= gap {
				n++
				array[j-gap], array[j] = iv, array[j-gap]
			}
		}
	}
	println("---------------", n)
}
