package main

import (
	"testing"
)

func TestStr(t *testing.T) {
	var tests = []struct {
		InPut, Close string
	}{
		{InPut: "http://vk.com and https://vk.com", Close: "http://****** and https://vk.com"},
		{InPut: "Ваша ссылка http://mail.ru", Close: "Ваша ссылка http://*******"},
		{InPut: "Сайт Пополам.РФ", Close: "Сайт Пополам.РФ"},
	}

	for _, tt := range tests {
		OutPut := mask(tt.InPut)
		if OutPut != tt.Close {
			t.Errorf("\ngot %s, want %s", OutPut, tt.Close)
		}
	}
}
