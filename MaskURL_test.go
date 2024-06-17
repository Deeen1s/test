package main

import (
	"testing"
)

func TestStr(t *testing.T) {
	var tests = []struct {
		name, InPut, Close string
	}{
		{name: "First: ", InPut: "http://vk.com and https://vk.com", Close: "http://****** and https://vk.com"},
		{name: "Second: ", InPut: "Ваша ссылка http://mail.ru", Close: "Ваша ссылка http://*******"},
		{name: "Third: ", InPut: "Сайт Пополам.РФ", Close: "Сайт Пополам.РФ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OutPut := mask(tt.InPut)
			if OutPut != tt.Close {
				t.Errorf("\ngot %s, want %s", OutPut, tt.Close)
			}
		})
	}
}
