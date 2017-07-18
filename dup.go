package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin,counts)
	}else{
		for _,arg:=range files{
			f,err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup: %v\n",err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}
}

func countLines(f *os.File,counts map[string]int){
	input:= bufio.NewScanner(f)
	for input.Scan(){
		fmt.Println(input.Text())
		counts[input.Text()]++
		fmt.Println(counts)
	}
}
