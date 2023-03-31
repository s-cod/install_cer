package main

import (
	"fmt"
	"os"
)

// import "github.com/gelsrc/go-charset"
func main() {

	filesIn, _ := os.ReadDir("./")
	var cer_file []string

	for i := range filesIn {
		name := filesIn[i].Name()
		if name[len(name)-3:] == "cer" {
			cer_file = append(cer_file, name)
		}
	}

	batFile, err := os.Create("install_cer.bat")
	if err != nil {
		fmt.Println("Ошибка создагтя bat файла")
	}
	defer func(f *os.File) {
		f.Close()
	}(batFile)

	batFile.WriteString("chcp 65001 >nul\r\n") //кодировка utf-8 для cmd

	// Для записи в файл в кодировке windows 1251
	// batFile.Write(charset.Cp1251RunesToBytes([]rune("chcp 1251 >nul\n")))
	// batFile.Write(charset.886 .Cp1251RunesToBytes([]rune("chcp 1251 >nul\n")))

	for i := range cer_file {
		// \r\n перенос строки в формате crlf
		f := fmt.Sprintf("csptest -ipsec -reg -autocont -mycert \"./%s\"\r\n", cer_file[i])
		// win := charset.Cp1251RunesToBytes([]rune(f))
		fmt.Println(f)
		// fmt.Println(win)
		batFile.WriteString(f)
	}
}
