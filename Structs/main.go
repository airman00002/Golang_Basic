package main

import "fmt"

type contactInfo struct { //todo <---
	email   string
	zipCode int
}

type person struct { //*กำหนด key และ value
	firstname   string
	lastname    string
	contactInfo //todo <---
}

func main() {

	// alex := person{firstname: "Alex", lastname: "Anderson"}
	// var alex person
	// alex.firstname = "Alex"
	// alex.lastname = "Anderson"

	// fmt.Println(alex)       //* ถ้าไม่มีค่า {}
	// fmt.Printf("%+v", alex) //* ถ้าไม่มีค่า {firstname: lastname:} *string ว่าง  แสดงทั้ง key value

	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 9400, //? ,
		}, //? missing ',' before newline in composite literal ต้องมี
	}

	// jimPointer := &jim //*เหมือนอางอิงตำแหน่งเดิม
	// jimPointer.updateName(("Jimmy_To_World"))
	jim.updateName(("Jimmy_To_World")) //*shoutcut ไปกำหนดที่  *person แค่นั้นสามารถอ้างอิง pointer ได้เหมือนกัน
	jim.print()                        //*<--- print
}

//* ต้องการ update ค่าบางอย่างใน struct  // *ชี้ไปที่ person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName //* operator ที่ชี้ไปยัง person
}

//*func ผูกกับ struct
func (p person) print() { //*<--- print
	fmt.Printf("%+v", p)

}
