package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// ./filehandling -fpath=/path-of-file/test.txt
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	// read whole file
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Printf("%T\t%v\n", data, data)
	fmt.Printf("%v\n", string(data))

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\t%v\n", f, f)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// read by chunks
	/*r := bufio.NewReader(f)
	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[:n]))
	} */

	// read by lines
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

}
