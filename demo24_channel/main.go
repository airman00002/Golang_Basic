package main

import "fmt"

func main() {

	ch := make(chan string, 2)

	ch <- "one" //*send
	fmt.Println("step one")
	fmt.Println(<-ch) //*การ recive ทำให้ buffer ว่าง
	ch <- "two"       //*send
	fmt.Println("step two")
	//* เมื่อไม่มี routine จะวิ่งแค่ใน main || ถ้าไม่มีจอง buffer ให้ channel จะถูก block ไม่สามารถวิ่งไปบรรทัดอื่น || อีกวิธีคือ goroutines demo25**
	//* 1. จอง buffer
	//*error: all goroutines are asleep - deadlock!
}
