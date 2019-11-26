package main

//import "fmt"
//import "io"
//
//type Human struct {
//	name string
//	age  int
//}
//
//type Human1 *Human
//
//func (Human1)say()  {//接受者不能为pointer类型或者接口类型
//	fmt.Println("say hi")
//}
//
//func main4579() {
//	h:=Human1{
//		name: "anko",
//		age:  20,
//	}
//	h.say()
//	(&h).say()
//}













//package main
//
//import (
//"fmt"
//)
//
//type notifier interface {
//	notify()
//}
//
//type user struct {
//	name string
//	email string
//}
//
//func (u *user) notify() {
//	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
//}
//
//func main() {
//	u := user{"Bill", "bill@email.com"}
//	//sendNotification(u)//报错，u必须为*user类型
//	sendNotification(&u)
//	u.notify()
//	(&u).notify()
//
//}
//// sendNotification 接受一个实现了 notifier 接口的值
//// 并发送通知
//func sendNotification(n notifier) {
//	fmt.Println("sendNotification----start")
//	n.notify()
//	fmt.Println("sendNotification----end")
//}