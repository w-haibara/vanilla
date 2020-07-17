package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var debug_mode bool = true

func TestEchoAPIHandler(t *testing.T) {
	msgs := []string{
		"wlfs6EBTDd2tuafo7UcnWUmsbE+ypJrpldbk0FjXVGIIZ7GW4U/nmHiq/LX86mS8VvMMhzK5qmVEwfB/L9tESdLjTpX9bqzQ5miATr42HuIafrotNRqrL9bSzl5+L6Hy6CWoprt2iHtMjzPn3QvWKEHTRAulfnXJoiBJsaVncYB5d8V+F9E+NsU3Ebg44WUZLWbYu/ox1HgdZeaYXyZlDTZTeuzoznfVCPu2IV5pCy/7Iykp5EjjrH3w9mBojFgbqYxJY5YXtQPzWUIqwugFy4EzSs8OicF+rRwD6wVV3RgvOakHeBGSE0XvshafTohkHH69D4bhUb0JFYhXCAJNHQ==",
	}

	for _, msg := range msgs {
		req, err := http.NewRequest("POST", "/enc/echo", strings.NewReader(msg))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(echoAPIHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got [%v] want [%v]", status, http.StatusOK)
		}

		expected := msg

		if body := rr.Body.String(); body != expected {
			t.Errorf("handler returned unexpected body: \ngot [%v] \nwant [%v]", body, expected)
		}
	}
}
