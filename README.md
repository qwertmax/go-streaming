# Streaming Basics

```go
func ingest() <-chan []string {
	out := make(chan []string)
	go func() {
		out <- []string{"aaaa", "bbb"}
		out <- []string{"cccccc", "dddddd"}
		out <- []string{"e", "fffff", "g"}
		close(out)
	}()
	return out
}

func process(concurrency int, in <-chan []string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	out := make(chan string)

	work := func() {
		for data := range in {
			for _, word := range data {
				// do any changes here
				word = word + " maxi"

				// pull out results
				out <- word
			}
		}
		wg.Done()

	}

	go func() {
		for i := 0; i < concurrency; i++ {
			go work()
		}

	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func store(in <-chan string) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for data := range in {
			fmt.Println(data)
		}
	}()
	return done
}

func main() {
	// stage 1 ingest data from source
	in := ingest()

	// stage 2 - process data
	reduced := process(4, in)

	// stage 3 - store
	<-store(reduced)
}

```
