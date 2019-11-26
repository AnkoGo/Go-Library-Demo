package main

import "fmt"

type ls []string

type ls_chid ls

func (*ls_chid) name() {
	fmt.Println("name....")
}

func main236237() {

	ls_chid_obj := ls_chid{"abc", "你好啊", "edf"}

	for _, v := range ls_chid_obj {
		fmt.Println(v)
	}
	ls_chid_obj.name()
	fmt.Printf("%T---", ls_chid_obj)
}
