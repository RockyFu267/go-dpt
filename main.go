package main

import (
	"fmt"
	"sync"
	"time"
)

var i int64
var i1 int64
var wg sync.WaitGroup
var mtx sync.Mutex
var intTMP int64
var pudTime time.Duration
var cusTime time.Duration

func main() {
	wg.Add(2)
	fmt.Println("hi")
	a := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890111111111234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890111111111234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901111111112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789011111111"
	fmt.Println(len(a))

	chTmp := make(chan string, 1100000)

	var strTMP string
	go timetick()
	go addByteChan(a, chTmp)
	go consumeByteCh(chTmp)

	fmt.Println(len(strTMP))
	wg.Wait()
	fmt.Println(len(chTmp), "-------")
	fmt.Println("生产时间: ", pudTime, "\n消费时间: ", cusTime)
}

func timetick() {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		i = i + 1
		fmt.Println(i, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		if i > 10 {
			break
		}
	}
}

func addByteChan(a string, chTmp chan string) {
	start := time.Now()
	defer wg.Done()
	for k := 0; k < 1000*1024; k++ {
		chTmp <- a
		if k%400 == 0 {
			fmt.Println(k)
		}
	}
	pudTime = time.Since(start)
	fmt.Println("pudTime =", pudTime)
}

func consumeByteCh(chTmp chan string) {
	start := time.Now()
	defer wg.Done()
	for {
		<-chTmp
		i1 = i1 + 1
		if i1%400 == 0 {
			fmt.Println(i1, "我在消费")
		}
		if i1 >= 1024000 {
			break
		}
	}
	cusTime = time.Since(start)
	fmt.Println("cusTime =", cusTime)
	fmt.Println(i1)
}

func addByte(a string, strTMP string) {
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

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}
