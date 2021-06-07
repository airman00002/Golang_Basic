package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//* func Writer
type logWriter struct{}

func main() {

	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	//*ความยุ่งยาก มีวิธีที่ง่ายขึ้น
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs)) //*บังคับ string

	//*io.Copy 2 argument -writer -reader(resp.body)
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs)) //*byte บังคับ string
	fmt.Println("Just wrote this many bytes : ", len(bs))

	//* int ที่เป็น len ของ bs
	return len(bs), nil
}
