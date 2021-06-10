package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1 Somthing")
	}
}
func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2 Somthing")
	}
}

func main() {
	go run1() //*แตก routines
	go run2() //*แตก routines
	time.Sleep(5 * time.Second)

}
