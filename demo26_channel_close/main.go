package main

import "fmt"

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	//*แก้ปัญหาไม่ต้องดักรอ sleep ของ routine
	//*แต่จะฟังข้อความจาก routine ว่าทำงานเสร็จแล้ว คือ * close channel

	close(c) //*ยกเลิก และทำลายทิ้ง
}

func main() {
	ch := make(chan int, 1)
	go routine1(ch, 10)

	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("No more Data")

	/*
		for true {
			value, ok := <-ch //*เช็ค channel ถ้าไม่มี close จะวน loop จน deadlock
			//*ok true channel เปิดอยู่ , false channel ปิด
			if !ok { //*ถ้า ok เป็น false
				fmt.Println("No more Data")
				break
			}
			fmt.Println(value)
		}
	*/
}
