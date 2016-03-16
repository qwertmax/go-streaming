package main

import (
	"fmt"
	"sync"
)

func load() <-chan []string {
	out := make(chan []string)

	go func() {
		// out <- []string{"111", "222", "333"}
		// out <- []string{"444", "555", "666"}
		// out <- []string{"777", "888", "999"}

		for i := 0; i < 10; i++ {
			out <- []string{fmt.Sprintf("%d", i)}
		}
		close(out)
	}()

	return out
}

func process(in <-chan []string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(4)

	work := func() {
		for val := range in {

		}
	}

}

func save(in <-chan []string) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for val := range in {
			// fmt.Printf("%#v\n", val[0])
			fmt.Printf("%#v\n", val[0])

			// for k, v := range val {
			// 	fmt.Printf("key %d: %s\n", k, v)
			// }
		}
	}()

	return done
}

func main() {
	in := load()

	<-save(in)
}
