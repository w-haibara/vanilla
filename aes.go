package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"net/http"
)

type Aes struct {
	enc cipher.Stream
	dec cipher.Stream
}

func NewAes(key []byte, iv []byte) Aes {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	return Aes{
		cipher.NewCFBEncrypter(c, iv),
		cipher.NewCFBDecrypter(c, iv),
	}
}

func (a Aes) Enc(enc_in []byte) []byte {
	enc_out := make([]byte, len(enc_in))
	a.enc.XORKeyStream(enc_out, enc_in)
	return enc_out
}

func (a Aes) Dec(dec_in []byte) []byte {
	dec_out := make([]byte, len(dec_in))
	a.dec.XORKeyStream(dec_out, dec_in)
	return dec_out
}

func encEchoAPIHandler(w http.ResponseWriter, r *http.Request) {
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key := []byte("1234567890123456")

	a := NewAes(key, iv)
	input := make([]byte, 256)
	r.Body.Read(input)
	msg := a.Dec(input)

	//TODO: some processes using 'msg'
	fmt.Printf("msg: %s\n", msg)

	output := a.Enc(msg)
	w.Write(output)
}
