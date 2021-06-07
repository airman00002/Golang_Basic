package main

import "fmt"

type shape interface {
	getArea() float64
}

//* สามเหลี่ยม
type triangle struct {
	height float64
	base   float64
}

//*พื้นที่
type square struct {
	sideLength float64
}

func main() {

	//*กำหนดค่าใน struct
	tr := &triangle{
		height: 3,
		base:   7,
	}

	sq := &square{
		sideLength: 4,
	}

	//* call func interfaces
	printArea(tr)
	printArea(sq)
	//**** Correct ******
}

//*เขียน func ที่รับ argument เป็น Interfaces
func printArea(s shape) {
	fmt.Println(s.getArea())
}

//*เขียน func เข้าไปใน triangle ,square //func ต้องเหมือนกับ interfaces
//* 1
func (tr triangle) getArea() float64 {
	return 0.5 * tr.height * tr.base
}

//*2
func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
