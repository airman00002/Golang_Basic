package main

import "fmt"

//*Interface
type bot interface { //*interface bot
	getGreeting() string //* มี func getGreeting return เป็น string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	// printGreeting(eb)
	// printGreeting(sb)
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) { //*สร้าง func ที่รับ interface เป็น argument ! || reuse func ได้
	fmt.Println(b.getGreeting()) //*ใช้ func ใน interface
}

//*ใช้ func ชื่อซ้ำกันไม่ได้เพราะ ไม่เป็น interfaces
// func printGreeting(eb englishBot) { //*<--
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) { //*<--
// 	fmt.Println(sb.getGreeting())
// }

//*ยัด func เข้าไปในตัวแปร
func (englishBot) getGreeting() string {
	//*VERY custom logic for generating an english greeting
	return "Hi there !"
}

//*ยัด func เข้าไปในตัวแปร
func (spanishBot) getGreeting() string {
	return "Hola !"
}
