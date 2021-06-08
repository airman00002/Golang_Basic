package main

import "fmt"

/*
	?Defer
	*func ที่ให้ทำหลังจากที่ code run เสร็จแล้ว
	*defer ไหนอยู่บนสุดจะทำหลังสุด เริ่มจากล่างขึ้นบน **
	*connect DB radis ....
	db := connect(DB) เชื่อมต่อ
	 defer db.close() ปิดการเชื่อมต่อ แต่ทำหลังสุด (defer*)
	insert()
	delete()

	?Panic
	*ให้ออกจาก program ไปเลย เมื่อใช้จะออกจาก program ตั้งแต่บรรทัดที่ใช้
	*ใช้กับอะไรที่รุนแรง เช่น connect Db
	db := connect()
	if db == error{
		panic("Connect Db error")
	}
	?Recover
	*มักจะใช้ร่วมกับ defer ทุกครั้ง
	*ซึ่งจะนำ case ของ panic มาพิจารณา
	*เช่น มี panic หลายตัว ..connect DB ไม่ได้ เปิด file ไม่ได้ ....
	defer func() {
		v := recover()
		if v == "Connect Db error" {
			fmt.Println("Error 1")
		}
		if v == "A3 error" {
			fmt.Println("Error 2")
		}
	}()

*/
func main() {

	defer a1()
	// panic("Connect Db error")
	a2()
	// panic("A3 error")
	a3()

}

func a1() {
	fmt.Println("A1")
}
func a2() {
	fmt.Println("A2")
}
func a3() {
	fmt.Println("A3")
}
