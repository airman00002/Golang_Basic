package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*sync.Mutex ทำงานกับ go routine
ในกรณีที่ต้องการเข้าถึงตัวแปรที่แชร์กันอยู่ ระหว่าง Go routine
ป้องกันไม่ให้ แทรกแซงในขณะที่ตัวแปรกำลัง Mutex อยู่ ( ถูกเปลี่ยนแปลงอยู่ )
*Mutex  =  ป้องกันการเปลี่ยนแปลงตัวแปรที่แชร์ระหว่าง routine โดยที่การเปลี่ยนแปลงยังไม่สมบูรณ์ โดยการ Lock และ Unlock
*/

// SafeCounter is safe to use concurrently.
//*1 สร้าง struct
type SafeCounter struct {
	v   map[string]int //*สร้าง map ใน struct
	mux sync.Mutex     //*ใครที่ต้องการเข้าถึง Struct นี้อย่างปลอดภัย สามารถสร้าง sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) { //*Inc เพิ่มค่า count
	c.mux.Lock() //*Lock ป้องกันไม่ให้ Go routine ตัวอื่นมาเปลี่ยนแปลง object ของ Struct จนกว่า operation นี้จะเสร็จสมบูรณ์
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++ //*ต้องทำอันนี้ให้เสร็จก่อน
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock() //*Lock ว่ากำลังจะดึงค่า ห้ามเขียนข้อมูลเข้าไป
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock() //*unlock ตอนท้ายสุดของ func
	return c.v[key]      //*ระหว่างนี้ return ค่ากลับไป
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") //*["key"] = .... มี 1000 routine //เพื่อป้องการการแทรกแซงระหว่าง อ่าน/เขียน โดยใช้ Mutex
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) //*เอาค่ามา Show
}
