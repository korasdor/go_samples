package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	filePath1 := path.Join(path.Dir(filename), "/tmp/dat1")

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filePath1, d1, 0644)
	check(err)

	filePath2 := path.Join(path.Dir(filename), "/tmp/dat2")

	f, err := os.Create(filePath2)
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
