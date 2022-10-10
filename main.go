package main

import (
	"fmt"
	"sync"
	"time"
)

var i int64
var wg sync.WaitGroup
var mtx sync.Mutex
var intTMP int64

func main() {
	wg.Add(3)
	fmt.Println("hi")
	a := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890111111111234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890111111111234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111"
	fmt.Println(len(a))

	var strTMP string
	go timetick()
	// // 400M
	// for k := 0; k < 4*1024; k++ {
	// 	strTMP = strTMP + a
	// 	if k%400 == 0 {
	// 		fmt.Println(k)
	// 	}
	// 	//fmt.Println(k)
	// }
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	// go addByte(a, strTMP)
	go addint()
	go addint()
	// for k := 0; k < 1000000; k++ {
	// 	intTMP = intTMP + 1
	// 	fmt.Println(intTMP)
	// }

	fmt.Println(len(strTMP))
	wg.Wait()

}

func timetick() {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		i = i + 1
		fmt.Println(i, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		if i > 20 {
			break
		}
	}
}

func addByte(a string, strTMP string) (res string) {
	defer wg.Done()
	for k := 0; k < 40*1024; k++ {
		mtx.Lock()
		strTMP = strTMP + a
		if k%400 == 0 {
			fmt.Println(k)
		}
		mtx.Unlock()
		//fmt.Println(k)
	}
	return res
}

func addint() {
	defer wg.Done()
	for k := 0; k < 500000; k++ {
		mtx.Lock()
		intTMP = intTMP + 1
		fmt.Println(intTMP)
		mtx.Unlock()
		//fmt.Println(k)
	}
}
