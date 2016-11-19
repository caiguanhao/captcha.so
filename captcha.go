package main

import (
	// #include <stdlib.h>
	"C"
	"bytes"
	"encoding/base64"
	"image/png"
	"unsafe"

	"github.com/dchest/captcha"
)

//export NewCaptcha
func NewCaptcha(identifier, _data *C.char, width, height C.int) *C.char {
	data := C.GoString(_data)
	var numbers []byte
	for _, c := range data {
		n := c - 48
		if 0 <= n && n <= 9 {
			numbers = append(numbers, byte(n))
		}
	}
	img := captcha.NewImage(C.GoString(identifier), numbers, int(width), int(height))
	var buf bytes.Buffer
	png.Encode(&buf, img)
	str := C.CString(base64.StdEncoding.EncodeToString(buf.Bytes()))
	defer C.free(unsafe.Pointer(str))
	return str
}

func main() {}
