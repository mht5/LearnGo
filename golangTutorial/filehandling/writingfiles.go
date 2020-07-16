package main

import (
	_ "flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func writeStringToAFile(f *os.File, s string) {
	l, err := f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeBytesToAFile(f *os.File, c []byte) {
	n2, err := f.Write(c)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeLineByLineToAFile(f *os.File, d []string) {
	var err error
	for _, v := range d {
		fmt.Fprintln(f, v)
		_, err := f.WriteString(v + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func appendToAFile() {
	f, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	_, err = f.WriteString(newLine + "\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

func produce(data chan<- int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data <-chan int, done chan<- bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func writeConcurrentlyToAFile() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

func main() {
	/* fptr := flag.String("fpath", "test1.txt", "file path to read from")
	flag.Parse()
	f, err := os.Create(*fptr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// writeStringToAFile(f, "Hello, World!")

	// d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	// writeBytesToAFile(f, d2)

	d := []string{"Welcome to the world of Go.", "Go is a compiled language.", "It is easy to learn Go."}
	writeLineByLineToAFile(f, d) */

	// appendToAFile()

	writeConcurrentlyToAFile()
}
