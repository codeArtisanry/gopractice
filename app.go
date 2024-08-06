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
