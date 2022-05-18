package main

import (
	"fmt"
	"sync"
	"time"
)

func dev07(channels ...<-chan interface{}) <-chan interface{} {
	merged := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))
	output := func(sc <-chan interface{}) {
		for sqr := range sc {
			merged <- sqr
		}
		wg.Done()
	}
	for _, optChan := range channels {
		go output(optChan)
	} 
	wg.Add(1)
	surveillance := func(sc <-chan interface{}) {
		for sqr := range sc {
			merged <- sqr
		}
		wg.Done()
	}
	for _, ch := range channels {
		go surveillance(ch)
	}
	wg.Wait()
	return merged
}

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or=dev07
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(9*time.Second),
		sig(8*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
