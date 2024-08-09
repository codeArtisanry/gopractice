package main

// func main()  {
// ch := make(chan int, 2) // buffered channel
// go func() {
// 	for i := 0; i <3; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }()

// time.Sleep(time.Second)
// for x := range ch {
// 	fmt.Println(x)
// }
// }

// output 0 1 2
// Explanation : In the above code, we have created a buffered channel ch of size 2. We have created a goroutine that sends 3 values to the channel ch. After sending all the values, we have closed the channel. We have used a for loop to receive the values from the channel ch. Since the channel is closed, the for loop will exit after receiving all the values from the channel ch.

// func main() {
// 	m := make(map[int]*int)

// 	for i := 0; i < 3; i++ {
// 		m[i] = &i
// 	}

// 	for k, v := range m {
// 		println(k, *v)
// 	}
// }

// output 0 2 1 3 3 3
// Explanation : In the above code, we have created a map m of type int to a pointer to int. We have added 3 key-value pairs to the map m. We have used a for loop to iterate over the map m. Since the value of i is updated in each iteration of the loop, all the values in the map m will point to the same memory location. Hence, when we print the values, we get the same value for all the keys.

// fibbonaci series in go
// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }
// Explanation : In the above code, we have defined a recursive function fib that calculates the nth Fibonacci number. The function takes an integer n as input and returns the nth Fibonacci number. If n is less than or equal to 1, the function returns n. Otherwise, it calculates the nth Fibonacci number by recursively calling the function with n-1 and n-2 as arguments and adding the results.

// can we the same code of fibbonaci series using memoization
// func fib(n int, memo map[int]int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	if val, ok := memo[n]; ok {
// 		return val
// 	}
// 	memo[n] = fib(n-1, memo) + fib(n-2, memo)
// 	return memo[n]
// }

// func memoizedFibonacci() func(int) int {
//     cache := make(map[int]int)
//     var fib func(int) int
//     fib = func(n int) int {
//         if val, found := cache[n]; found {
//             return val
//         }
//         if n <= 1 {
//             return n
//         }
//         result := fib(n-1) + fib(n-2)
//         cache[n] = result
//         return result
//     }
//     return fib
// }

// func accumulator() func(i int) int {
//     sum := 0
//     return func(n int) int {
//         sum += n
//         return sum
//     }
// }

// func stringer() func(s string) string {
//     summer := ""
//     return func(b string) string {
//         summer += " "
//         summer += b
//         return summer
//     }
// }

// func largest(i []int) int {
// 	max := i[0] // assuming 0th index is big

// 	for _, v := range i {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }
