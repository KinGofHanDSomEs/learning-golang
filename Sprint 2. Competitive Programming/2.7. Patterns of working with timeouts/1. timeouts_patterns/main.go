package main  

import (  
	"fmt"  
	"time"  
)  

func isPrime(num int) bool {  
	if num <= 1 {  
		return false  
	}  
	for i := 2; i*i <= num; i++ {  
		if num%i == 0 {  
			return false  
		}  
	}  
	return true  
}  

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {  
	for num := 2; num <= N; num++ {  
		if isPrime(num) {  
			prime_nums <- num 
			time.Sleep(10 * time.Millisecond)   
		}  

		select {  
		case <-stop:  
			close(prime_nums)
			return  
		default:  
		}  
	}  
	close(prime_nums)
}  

func main() {  
	N := 100000 // Заданное число N  
	stop := make(chan struct{})  
	prime_nums := make(chan int)  

	go GeneratePrimeNumbers(stop, prime_nums, N)  

	// Вывод простых чисел  
	for prime := range prime_nums {  
		fmt.Println(prime)  
	}  

	// Что-то может потребоваться в будущем для остановки (пример через 1 сек):  
	time.Sleep(1 * time.Second)  
	close(stop) // Отправка сигнала остановки  
}