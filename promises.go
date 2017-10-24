package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	urls := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	seen := make(map[int]bool)

	t0 := time.Now()

	for {

		ch := make(chan []int, len(urls))

		var wg sync.WaitGroup

		for _, v := range urls {

			wg.Add(1)

			go func(v int) {
				defer wg.Done()
				ch <- doit(v)
			}(v)

		}

		wg.Wait()
		close(ch)

		for v := range ch {
			urls = urls[1:]

			for _, j := range v {

				if seen[j] != true {
					urls = append(urls, j)
				}
				seen[j] = true
			}

			//	fmt.Printf("%+v\n", urls)

		}

		if len(urls) == 0 {
			break
		}

	}

	td := time.Now().Sub(t0)
	fmt.Printf("%+v", td)

}

func doit(url int) []int {
	time.Sleep(2 * time.Second)

	// Seed the number random generator.
	rand.Seed(time.Now().UnixNano())

	// Choose a random number from 0 to 9.
	rn := rand.Intn(9)

	fmt.Printf("doit %d %d\n", url, rn)

	return []int{100, 101}

}
