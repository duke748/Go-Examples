package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func cheese() {
	for i := 0; i < 11; i++ {
		fmt.Println("cheese", i)
	}
	wg.Done()
}

func burger() {

	for i := 11; i < 21; i++ {
		fmt.Println("burger", i)
	}

}

func main() {

	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal("Fuxxake", err)
		continue
	}
	fmt.Println(f)

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUS\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go cheese()
	burger()
	wg.Wait()

	fmt.Println("CPUS\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}
