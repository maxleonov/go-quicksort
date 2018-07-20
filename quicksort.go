package main

func Quicksort(a []interface{}, less func(a interface{}, b interface{}) bool) []interface{} {
	// todo shuffle
	doSort(a, 0, len(a)-1, less)
	return a
}

func doSort(a []interface{}, lo int, hi int, less func(a interface{}, b interface{}) bool) {
	if hi <= lo {
		return
	}

	var j int = partition(a, lo, hi, less)
	doSort(a, lo, j-1, less)
	doSort(a, j+1, hi, less)
}

func partition(a []interface{}, lo int, hi int, less func(a interface{}, b interface{}) bool) int {
	var i int = lo + 1
	var j int = hi
	var v = a[lo]

	for {
		for less(a[i], v) {
			if i == hi {
				break
			}
			i++
		}
		for less(v, a[j]) {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		exch(a, i, j)
	}
	exch(a, lo, j)

	return j
}

func exch(a []interface{}, i int, j int) {
	var tmp interface{}

	tmp = a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {
}
