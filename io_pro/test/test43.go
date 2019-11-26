package main
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main345667() {

	ls:=[]int{1,2,3,4,5}
	ch := make(chan int)

	//取值
	go func (){
		ret:=0
		for {
			n := <-ch
			ret+=n//事实上你的这些操作都可以在存值的时候做，而不是在这里做，整体表现出你的需求不明确
			fmt.Printf("%v接收成功,和为%v\n",n, ret)
			last_index:=len(ls)-1

			if n == ls[last_index]*ls[last_index]{
				break
			}
		}
		defer wg.Done()
	}()
	wg.Add(1)

	//存值
	go func() {
		for _, v := range ls {//如果是引用类型而使用&ls的话会报错，假如是值类型则不会
			ch <- v*v
		}
		defer wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	fmt.Println("发送成功")
}


