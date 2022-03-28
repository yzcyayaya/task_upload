package main

import (
	"fmt"
	"sync"
)

//var flag int

//func main() {
//	flag = 1
//	bools := make(chan bool,2)
//	bools<-true
//	go func() {
//		for flag<=100 {
//			if <-bools{
//				fmt.Println("g1====",flag)
//				flag=flag+1
//			}
//			bools<-true
//		}
//
//	}()
//	go func() {
//		for flag<=100 {
//			if <-bools{
//				fmt.Println("g2====",flag)
//				flag=flag+1
//			}
//			bools<-true
//		}
//
//	}()
//	go func() {
//		for flag<=100 {
//			if <-bools{
//				fmt.Println("g3====",flag)
//				flag=flag+1
//			}
//			bools<-true
//		}
//
//	}()
//	go func() {
//		for flag<=100 {
//			if <-bools{
//				fmt.Println("g4====",flag)
//				flag=flag+1
//			}
//			bools<-true
//		}
//
//	}()
//	go func() {
//		for flag<=100 {
//			if <-bools{
//				fmt.Println("g5====",flag)
//				flag=flag+1
//			}
//			bools<-true
//		}
//	}()
//	time.Sleep(1*time.Second)
//}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	group := sync.WaitGroup{}
	group.Add(3)
	go C(ch2, &group)
	go B(ch1, ch2, &group)
	go A(ch1, &group)
	group.Wait()
}
func A(ch1 chan bool, wg *sync.WaitGroup) {
	fmt.Println("A")
	ch1 <- true
	wg.Done()
}
func B(ch1 chan bool, ch2 chan bool, wg *sync.WaitGroup) {
	<-ch1
	fmt.Println("B")
	ch2 <- true
	wg.Done()
}
func C(ch2 chan bool, wg *sync.WaitGroup) {
	<-ch2
	fmt.Println("C")
	wg.Done()
}
