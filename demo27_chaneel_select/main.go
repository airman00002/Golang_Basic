package main

import "fmt"

/*
*select
ทำให้ go routine สามารถรอกระทำการสื่อสารแบบ mutiple ได้
select จะ block จนกว่า case เหล่านั้นจะถูก run
*/

func task(c, quit chan string) {
	for {
		//*สุ่มว่ารอบนี้ควร run channel ไหน
		//*ถ้าไม่มีข้อมูลยิงเข้ามาก้จะทำ cas อื่น
		//! ถ้าไม่มีเลย จะถูก block ให้รอเฉยๆ **deaclock
		select {
		case c <- "Running ...":
		case <-quit:
			fmt.Println("quit")
			return //*return ออกจาก while loop****
		}
	}
}

func main() {
	//*สร้าง channel มากกว่า 1 ตัว
	//* select คือ ดึงข้อความจาก channel ไหนก่อน
	//*ในที่นี้ตัวนึงแลกเปลี่ยนข้อมูล อีกตัวส่งข้อมูลให้หยุดการทำงาน
	c := make(chan string)
	quit := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //*send
		}
		quit <- "Go !" //*recive อะไรเข้าไปก้ได้
	}()
	task(c, quit)
}
