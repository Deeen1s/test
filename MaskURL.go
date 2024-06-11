package main

import (
	"bufio"
	"fmt"
	"os"
)

func mask(messege string) string { //функция

	sliceWeb := []byte(messege) //пустой байтовый слайс для сайта
	var sliceHttp []byte        //пустой байтовый слайс для http
	httpStr := "http://"
	sliceWeb = []byte(messege)  //перевод ссылки в байты
	sliceHttp = []byte(httpStr) //перевод  в байты http://
	lenHttp := len(sliceHttp)   //длина http://
	lenMessege := len(messege)  //длина входных данных в программе

	for i := 0; i < lenMessege; {
		if i <= lenMessege-lenHttp && string(sliceWeb[i:i+lenHttp]) == httpStr { //проверяем элементы на схожесть с http://
			i += lenHttp
			for i < lenMessege && sliceWeb[i] != ' ' { // маскировка ссылки до пробела
				sliceWeb[i] = '*'
				i++
			}
		} else {
			i++
		}
	}
	return string(sliceWeb)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin) // создаем  bufio.Scanner
	fmt.Print("Вставьте ссылку: ")        // просим вписать сайт
	_ = scanner.Scan()                    // вписываем сайт
	webSite := scanner.Text()             // создаем через буфер

	fmt.Println("Шифрованный сайт или сайты:", mask(webSite)) // вывод шифрованного сайта

}
