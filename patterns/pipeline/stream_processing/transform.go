package main

// Applies multiple transformations to the data.
func transform(in <-chan int, transformers ...func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			for _, transform := range transformers {
				n = transform(n)
			}
			out <- n
		}
		close(out)
	}()
	return out
}
