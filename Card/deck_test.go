package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //* Upper case //*ถูกเรียกโดย go test runner
	d := newDeck()

	if len(d) != 16 { //* %v = ใส่ตัวแปร ที่เป็น argument เข้าไป
		t.Errorf("Expected deck lenght of 16 , but got %v ", len(d))
	}

	//* ถ้าข้อมูลตำแหน่งแรก ไม่ถูกต้อง
	if d[0] != "Ace of Speades" {
		t.Errorf(("Expected first card of Ace of Speades , but got %v"), d[0])
	}
	//* ถ้าข้อมูลตำแหน่งสุดท้าย ไม่ถูกต้อง
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf(("Expected last card of Four of Clubs , but got %v"), d[0])
	}
}

//* บันทึกไฟล์ โหลดไฟล์ และลบ
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	//*ลบไฟล์เดิมออก
	os.Remove("_decktesting")
	//*สร้าง deck และ save
	deck := newDeck()
	deck.savetoFile("_decktesting")

	//*โหลดไฟล์
	loadedDeck := newDeckFromFile("_decktesting")

	//*check ขนาดไฟล์
	if len(loadedDeck) != 12 {
		t.Errorf("Expected 16 cards in deck , got %v ", len(loadedDeck))
	}

}
