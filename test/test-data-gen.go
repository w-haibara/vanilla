package main

import (
	"encoding/base64"
	"fmt"

	"github.com/w-haibara/vanilla"
)

func gen(key, iv []byte, msgs []string) {
	for _, msg := range msgs {
		aes := vanilla.NewAES(key, iv)
		encrypted := base64.StdEncoding.EncodeToString(aes.Enc([]byte(msg)))
		fmt.Println(msg, encrypted)
	}
}

func main() {
	key := []byte("1234567890123456")
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	msgs := []string{"Hello", "Hello Worls!", "Who are you?", "Golang", "Goodbye"}

	gen(key, iv, msgs)
}
