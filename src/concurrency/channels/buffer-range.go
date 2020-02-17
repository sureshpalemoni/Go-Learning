/* Using buffer channel and range, 
we can read the data from the closed channel
*/

package channels

import "fmt"

func  BufferRange()  {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	for elem := range c {
		fmt.Println(elem)
	}
}

