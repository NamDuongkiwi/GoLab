package handle

func GetMinOfArray(array []int)(m int){
	for i, value := range array {
		if i==0 || value < m {
			m = value
		}
	}
	return
}

func GetMaxOfArray(array []int)(m int){
	for i, value := range array {
		if i==0 || value > m {
			m = value
		}
	}
	return
}

func AvgOfArray(array []int) float64 {
	var avg float64
	for _, value :=  range array{
		avg += float64(value)
	}
	return avg / float64(len(array))
}

func BubbleSort(array []int) []int{
	for i:= 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++{
			if array[i] > array[j]{
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	first := MergeSort(array[:len(array)/2])
	second := MergeSort(array[len(array)/2:])
	return merge(first, second)
}

func merge(first []int, second []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	return final
}

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1

	pivot := len(array)/2

	array[pivot], array[right] = array[right], array[pivot]

	for i, _ := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}

func checkPrimeINArray(array []int)  []int{
	var result []int
	for _, value := range array{
		if isPrime(value){
			result = append(result, value)
		}
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n/2 + 1 ; i ++{
		if n % i == 0{
			return false
		}
	}
	return true
}


