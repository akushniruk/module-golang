package goroutines

func Process(input chan string) chan string {
	out := make(chan string)
	res := make(chan bool)

	go func() {
		out <- "(" + <-input + ")"
		res <- true
	}()

	go func() {
		<-res
		close(out)
	}()

	return out
}
