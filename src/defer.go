package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	//достаем путь до этого файла
	_, filename, _, _ := runtime.Caller(0)
	//склеиваем полученный пути
	f := createFile(path.Join(path.Dir(filename), "/tmp/defer.txt"))
	//сработает на выходе функции
	defer closeFile(f)
	//пишем в файл
	writeFile(f)
}
