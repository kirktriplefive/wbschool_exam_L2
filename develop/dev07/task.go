package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} { // функция, создающая канал для чтения
		c := make(chan interface{}) // канал для закрытия sig канала
		go func() {                 // в горутине
			defer close(c)             // закрываем канал с
			time.Sleep(after)         
			fmt.Println(after, "gone") 
		}()
		return c
	}
   
	start := time.Now() 
	<-or(               // вызываем функцию, которая будет принимать в себя каналы, созданные через sig
		sig(1*time.Second), // создаем канал и закрываем его через секунду
		sig(3*time.Second), // создаем канал и закрываем его через три секунды
		sig(7*time.Second), 
		sig(10*time.Second), 
	)
   
	fmt.Printf("return after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup         
		output := make(chan interface{}) // создаем single-канал
	
		for i := range channels { 
			wg.Add(1) // значения каждого канала будут вычитываться в своих горутинах
			go func(input <-chan interface{}) {
				for val := range input { // итерируемся по значениям канала
					output <- val // пишем значение в единый канал
				}
				wg.Done() 
			}(channels[i])
		}
		wg.Wait() 
		close(output) //закрываем канал
		return output // возвращаем канал
}