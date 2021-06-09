package main

import "fmt"

func routines1(ch chan string, payload string) {
	ch <- payload
	fmt.Println("step one")
	ch <- payload
	fmt.Println("step two")
	ch <- payload
	fmt.Println("step three")

}

//*deadlock ไป block การทำงานของ main วิธีแก้มี 2 วิธี ..buffer ..go routines ย่อย
func main() {

	ch := make(chan string, 2) //*ถ้าจองน้อยเกินไปจะ block การทำงานของ routines ย่อย
	go routines1(ch, "one")
	// go routines1(ch, "two")
	// go routines1(ch, "three")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

}
