/*

*************Send and Receive Data From a Channel*****************
In Go language, channel work with two principal operations one is sending and another one is receiving, both the operations collectively known as communication. And the direction of <- operator indicates whether the data is received or send. In the channel, the send and receive operation block until another side is not ready by default. It allows goroutine to synchronize with each other without explicit locks or condition variables.

==>Send operation: The send operation is used to send data from one goroutine to another goroutine with the help of a channel.
Values like int, float64, and bool can safe and easy to send through a channel because they are copied so there is no risk of accidental concurrent access of the same value. Similarly, strings are also safe to transfer because they are immutable. But for sending pointers or reference like a slice, map, etc. through a channel are not safe because the value of pointers or reference may change by sending goroutine or by the receiving goroutine at the same time and the result is unpredicted. So, when you use pointers or references in the channel you must make sure that they can only access by the one goroutine at a time.
Mychannel <- element
The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.

==>Receive operation: The receive operation is used to receive the data sent by the send operator.
element := <-Mychannel
The above statement indicates that the element receives data from the channel(Mychannel).
If the result of the received statement is not going to use is also a valid statement.
You can also write a receive statement as:<-Mychannel
*/

package main 

import (
	"fmt"
)


func main () {

	c := make(chan int)
	cr := make (<-chan int) //receive
	cs := make(chan int) //send

	fmt.Println("---------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cr)
	fmt.Printf("cr\t%T\n", cs)

}
// func main (){
// 	c := make(chan int)

// 	c <- 42
// 	c <-43

// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// 	fmt.Println("---------")
// 	fmt.Printf("%T\n", c)




