package main
import (
	"fmt"
	"time"
)
func main45789() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("." )
			time.Sleep(5e7)//2个5e7=1e8
		}
	}
}