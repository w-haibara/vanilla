package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

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

func (a AES) DecReader(r io.Reader) io.Reader {
	buf := new(bytes.Buffer)
	io.Copy(buf, r)
	dec_in := buf.Bytes()
	dec_out := make([]byte, len(dec_in))
	a.dec.XORKeyStream(dec_out, dec_in)
	return bytes.NewReader(dec_out)
}

func (a AES) Dec(dec_in []byte) []byte {
	dec_out := make([]byte, len(dec_in))
	a.dec.XORKeyStream(dec_out, dec_in)
	return dec_out
}

type SecureWriter struct {
	origWriter http.ResponseWriter
	encWriter  *bufio.Writer
}

func (s *SecureWriter) WriteHeader(rc int) {
	s.origWriter.WriteHeader(rc)
}

func (s *SecureWriter) Write(p []byte) (int, error) {
	return s.encWriter.Write(p)
}

func (s *SecureWriter) Header() http.Header {
	return s.origWriter.Header()
}

func CryptoHandler(fn http.HandlerFunc) http.HandlerFunc {
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key := []byte("1234567890123456")
	return func(w http.ResponseWriter, r *http.Request) {
		aes := NewAES(key, iv)

		decoder := base64.NewDecoder(base64.StdEncoding, r.Body)
		r, err := http.NewRequest(r.Method, r.URL.String(), aes.DecReader(decoder))
		if err != nil {
			log.Print("new request error: ", err)
			return
		}

		wb := bytes.NewBuffer([]byte{})
		sw := &SecureWriter{origWriter: w, encWriter: bufio.NewWriter(wb)}
		fn(sw, r)
		sw.encWriter.Flush()
		encoded := base64.StdEncoding.EncodeToString(aes.Enc(wb.Bytes()))
		io.WriteString(w, encoded)
	}
}
