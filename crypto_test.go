package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var debug_mode bool = true

func TestEchoAPIHandler(t *testing.T) {
	msg := []byte("hello")
	block_size := 256
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key := []byte("1234567890123456")

	aes := NewAES(key, iv)
	block := make([]byte, block_size)
	for i := range msg {
		block[i] = msg[i]
	}

	encrypted := aes.Enc(block)
	encoded := base64.StdEncoding.EncodeToString(encrypted)

	if debug_mode {
		fmt.Printf("%s", hex.Dump(encrypted))
		fmt.Println(encoded)
	}

	req, err := http.NewRequest("POST", "/enc/echo", strings.NewReader(encoded))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(echoAPIHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got [%v] want [%v]", status, http.StatusOK)
	}

	decoded := make([]byte, block_size)
	_, err = base64.StdEncoding.Decode(decoded, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	expected := block
	if decrypted := aes.Dec(decoded); !bytes.Equal(decrypted, expected) {
		t.Errorf("handler returned unexpected body: \ngot [%v] \nwant [%v]", decrypted, expected)
	}
}
