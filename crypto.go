package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

var debud_mode bool = true

type AES struct {
	enc cipher.Stream
	dec cipher.Stream
}

func NewAES(key []byte, iv []byte) AES {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	return AES{
		cipher.NewCFBEncrypter(c, iv),
		cipher.NewCFBDecrypter(c, iv),
	}
}

func (a AES) Enc(enc_in []byte) []byte {
	enc_out := make([]byte, len(enc_in))
	a.enc.XORKeyStream(enc_out, enc_in)
	return enc_out
}

func (a AES) Dec(dec_in []byte) []byte {
	dec_out := make([]byte, len(dec_in))
	a.dec.XORKeyStream(dec_out, dec_in)
	return dec_out
}

type SecureWriter struct {
	origWriter http.ResponseWriter
	encWriter  func([]byte) (int, error)
}

func (s *SecureWriter) WriteHeader(rc int) {
	s.origWriter.WriteHeader(rc)
}

func (s *SecureWriter) Write(p []byte) (int, error) {
	return s.encWriter(p)
}

func (s *SecureWriter) Header() http.Header {
	return s.origWriter.Header()
}

func _crypto(fn http.HandlerFunc) http.HandlerFunc {
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key := []byte("1234567890123456")
	aes := NewAES(key, iv)
	buf := make([]byte, 256)

	return func(w http.ResponseWriter, r *http.Request) {
		var cipher []byte = nil
		for {
			_, err := r.Body.Read(buf)
			for _, v := range buf {
				cipher = append(cipher, v)
			}
			if err != nil {
				log.Print("read HTTP request body : ", err)
				break
			}
		}

		var decoded []byte = nil
		_, err := base64.StdEncoding.Decode(decoded, cipher)
		if err != nil {
			log.Print("decode error:", err)
			return
		}

		r, err = http.NewRequest(r.Method, r.URL.String(), bytes.NewReader(aes.Dec(decoded)))
		if err != nil {
			log.Print("new request error: ", err)
			return
		}

		/*
			sw := &SecureWriter{origWriter: w, encWriter: func(p []byte) (int, error) {

				return 0, nil
			}}

			fn(sw, r)
		*/

		fn(w, r)

		//w.Write(strings.NewReader(base64.StdEncoding.EncodeToString(aes.Enc(cipher))))
	}
}

func crypto(fn http.HandlerFunc) http.HandlerFunc {
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key := []byte("1234567890123456")
	return func(w http.ResponseWriter, r *http.Request) {
		aes := NewAES(key, iv)

		buf := make([]byte, 344)
		_, err := r.Body.Read(buf)
		if err != nil {
			log.Print("read HTTP request body : ", err)
		}

		decoded := make([]byte, 256)
		_, err = base64.StdEncoding.Decode(decoded, buf)
		if err != nil {
			log.Print("decode error:", err)
			return
		}
		decrypted := aes.Dec(decoded)

		if debud_mode {
			fmt.Printf("decoded\n%s", hex.Dump(decoded))
			fmt.Printf("decrypted\n%s", hex.Dump(decrypted))
		}

		r, err = http.NewRequest(r.Method, r.URL.String(), bytes.NewBuffer(decrypted))
		if err != nil {
			log.Print("new request error: ", err)
			return
		}

		sw := &SecureWriter{origWriter: w, encWriter: func(p []byte) (int, error) {

			return 0, nil
		}}

		fn(sw, r)
	}
}
