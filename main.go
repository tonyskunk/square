/*2. Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	square := func(n int) {
		defer wg.Done()
		fmt.Printf("Квадрат числа %d равен %d\n", n, n*n)
	}
	for _, number := range numbers {
		wg.Add(1)
		go square(number)
	}
	wg.Wait()
}
