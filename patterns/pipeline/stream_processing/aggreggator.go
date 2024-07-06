package main

// Sums up all numbers and calculates the average.
func aggregate(in <-chan int) (sum int, avg float64) {
	count := 0
	for n := range in {
		sum += n
		count++
	}
	if count > 0 {
		avg = float64(sum) / float64(count)
	}
	return
}
