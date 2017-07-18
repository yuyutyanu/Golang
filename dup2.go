package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main()  {
	counts := make(map[string]int)
	for _,filename:= range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
			continue
		}
		for _,line:= range strings.Split(string(data),"\n"){
			counts[line]++
		}
		for line,n := range counts{
			if  n > 1{
				fmt.Printf("%s:%d行の重複:%s\n",filename,n,line)
			}
		}
	}
}
