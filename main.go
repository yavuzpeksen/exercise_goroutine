package main 

import (
	"fmt"
	"time"
)

func f(n int){
	for i:=0; i<n; i++{
		fmt.Println(n," i:", i)
		
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func main() {
	for i:=5; i< 21; i+=5{
		go f(i)
	} 
	var input string
	fmt.Scanln(&input)
}

