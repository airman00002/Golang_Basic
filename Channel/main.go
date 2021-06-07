package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.fackbook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	//*กรณี พิม channel ออกแค่ 1 ครั้งจะได้แค่ goole เพราะค่าอื่นใน routine ยังทำไม่เสร็จ(ระหว่างทำ routine main sleepping)  //*อะไรเสร็จก่อนออกมาก่อน asyn
	// fmt.Println(<-c) //*รองรับข้อความจาก channel 1 ครั้ง
	//*การวน for ก้ยังเป็นการรอให้แต่ละรอบส่งค่ามาให้ channel **
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//*การ ping ซ้ำๆ  เรียกซ้ำๆ
	// for { //*1.string 2.channel
	// 	go checkLink(<-c, c)
	// }

	//*ไวยากรณ์ loop ทางเลือก //*ทำให้ดูเข้าใจขึ้น แต่ ping ตลอดเวลาอยู่ดี
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //*ส่งผ่าน l เป็น argument ของ func
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be Down !")
		// c <- "Might be down I think!"
		c <- link
		return
	}
	fmt.Println(link, "is Up !")
	// c <- "Yep its up"
	c <- link
}
