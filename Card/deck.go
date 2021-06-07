package main //*file ที่อยู่ใน package จะต้องประกาศชื่อด้านบนเสมอ เขียนใน package นี้
//*รวมประเภทของ deck เข้าด้วยกัน
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//*create a new type of deck
//* witch is a clice of string

type deck []string //*กำหนดประเภทใหม่ใน app

func newDeck() deck {
	cards := deck{} //* ประเภท deck

	cardSuits := []string{"Speades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//*จะสร้าง fnใหม่ ที่เป้น deck และเมื่อเรียก fn จะพิมแต่ละ deck
//*การวนซ้ำ card
//* d อ้างอิง deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//*เปลี่ยน []string ให้เป็นแค่ string
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //*return string ก้อนเดียว
}

func (d deck) savetoFile(fileName string) error { //*save ไว้ใน hard drive //return error ได้
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) //*สามารถอ่านและเขียน file ได้
}

//*  อ่าน file ใน hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //*อ่านไฟล์ ผลลัพธ์ 2แบบ ได้ไฟล์ , error
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error : ", err)
		os.Exit(1) //*ถ้ามีอะไรผิดพลาดต้องออกจาก Program
	}

	s := strings.Split(string(bs), ",") //* s : slice of strings
	return deck(s)
}

//*สลับตำแหน่ง cards //*ไม่มี return
func (d deck) shuffle() { //*d คือสำเนาของ deck
	//*เมื่อเริ่ม  func จะสร้าง source ใหม่ จากนั้นใช้ source สำหรับสร้างเลลขสุ่ม
	source := rand.NewSource(time.Now().UnixNano()) //* หาค่า source ใหม่จาก time 64 bit *UnixNano( 64 )
	r := rand.New(source)                           //*ได้ค่าเริ่มต้นใหม่ ไม่ซ้ำกัน  //* r * rand (int)

	for i := range d {
		//*random ค่า
		newPosition := r.Intn(len(d) - 1)
		//*สลับค่า
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//* Even Odd
func even_odd() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "Even")
		} else {
			fmt.Println(num, "Odd")
		}
	}
}
