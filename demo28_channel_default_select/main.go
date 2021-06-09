package main

import (
	"fmt"
	"time"
)

/*
*default
ถูกเรียกรณีที่เรารอ operation แต่ยังไม่มีงานของ channel เกิดขึ้นเลย
*/
func task(c, quit chan string) {
	for {

		select {
		case c <- "Running ...":
		case <-quit:
			fmt.Println("quit")
			return //*return ออกจาก while loop****
		default: //*ถ้าไม่มี case เกิดขึ้นเลย***
			fmt.Println("waiting....")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {

	c := make(chan string)
	quit := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //*send
		}
		// quit <- "Go !" //*recive อะไรเข้าไปก้ได้
	}()
	task(c, quit)
}
