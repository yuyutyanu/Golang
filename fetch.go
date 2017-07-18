package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	for _,url:= range os.Args[1:]{
		res,err:= http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
		b,err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		if	err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s : %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}
