package qsort

func partition(values []int, p, r int) {
	x := values[r]
	i := p - 1
	for j := p; j < r; j++ {
		if values[j] <= x {
			i++
			values[i], values[j] = values[j], values[i]
		}
	}
	values[i+1], values[r] = values[r], values[i+1]
	q := i + 1
	if q-p > 1 {
		partition(values, p, q-1)
	}
	if r-q > 1 {
		partition(values, q+1, r)
	}
}

func QuickSort2(values []int) {
	if len(values) > 0 {
		partition(values, 0, len(values)-1)
	}
}
