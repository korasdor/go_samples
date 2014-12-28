package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//с этим методом подхода к счетчиками надо разобраться, не совсем все понятно
func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			//метод получает как первый параметр указатель значения, второй параметр количество
			atomic.AddUint64(&ops, 1)

			runtime.Gosched()
		}()
	}

	//стопаем секунду что бы все утрамбовалось
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
