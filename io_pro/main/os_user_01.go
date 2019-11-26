package main

import (
	"fmt"
	"os/user"
)

func main236785() {
	user111,err:=user.Current()
	if err != nil{
		fmt.Println(err)
	}
	//输出结果：&{S-1-5-21-1198569543-3519917723-292291289-500 S-1-5-21-1198569543-3519917723-292291289-513 WIN-5EP8LP6GHN7\Administrator  C:\Users\Administrator}
	fmt.Println(user111)
	fmt.Println(user111.Name)//为空字符串
	fmt.Println(user111.Gid)//S-1-5-21-1198569543-3519917723-292291289-513
	fmt.Println(user111.Uid)//S-1-5-21-1198569543-3519917723-292291289-500
	fmt.Println(user111.Username)//WIN-5EP8LP6GHN7\Administrator
	fmt.Println(user111.GroupIds())//[S-1-5-32-544 S-1-5-21-1198569543-3519917723-292291289-513] <nil>
	fmt.Println(user111.HomeDir)//C:\Users\Administrator

	fmt.Println("---------------------")
	user222,err222:=user.Lookup(`WIN-5EP8LP6GHN7\Administrator`)
	if err222 != nil{
		fmt.Println(err222)
	}
	//&{S-1-5-21-1198569543-3519917723-292291289-500 S-1-5-21-1198569543-3519917723-292291289-513 WIN-5EP8LP6GHN7\Administrator  C:\Users\Administrator}
	fmt.Println(user222)

	fmt.Println("---------------------")
	user333,err333:=user.LookupId("S-1-5-21-1198569543-3519917723-292291289-500")
	if err333 != nil{
		fmt.Println(err333)
	}
	//&{S-1-5-21-1198569543-3519917723-292291289-500 S-1-5-21-1198569543-3519917723-292291289-513 WIN-5EP8LP6GHN7\Administrator  C:\Users\Administrator}
	fmt.Println(user333)


	fmt.Println("---------------------")
	user444,err444:=user.LookupGroupId("S-1-5-21-1198569543-3519917723-292291289-513")//注意这个返回的不是*user，而是*Group
	if err444 != nil{
		fmt.Println(err444)
	}

	fmt.Println(user444.Name)//None
	fmt.Println(user444.Gid)//S-1-5-21-1198569543-3519917723-292291289-513



}