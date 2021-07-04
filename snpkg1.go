package main

import ("fmt"
		"os"
)

func main() {
	show := func(key string){
		val,ok := os.LookupEnv(key)
		if !ok {
			fmt.Println("not set")
		}else{
			fmt.Println(key+":"+val)
		}
	}
	//os.Setenv("TEST","val")
	show("GOPATH")
	fmt.Println("Hello")
}
