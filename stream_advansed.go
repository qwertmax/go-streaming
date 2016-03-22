package main

import (
	"fmt"
	"sync"
)

func load() <-chan []string {
	out := make(chan []string)

	go func() {
		for i := 0; i < 10; i++ {
			out <- []string{fmt.Sprintf("%d", i)}
		}
		close(out)
	}()

	return out
}

func process(in <-chan []string, data <-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(4)

	out := make(chan string)

	work := func() {
		for str := range in {
			for _, val := range str {
				val = val + "!" + <-data
				out <- val
			}
		}
		wg.Done()
	}

	go func() {
		for i := 0; i < 4; i++ {
			go work()
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func save(in <-chan string) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for val := range in {
			fmt.Printf("%#v\n", val)
		}
	}()

	return done
}

func data_processing() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for i := 0; i < 12; i++ {
			out <- fmt.Sprintf("%d", i)
		}
	}()

	return out
}

func main() {
	in := load()
	data := data_processing()

	out := process(in, data)

	<-save(out)
}
