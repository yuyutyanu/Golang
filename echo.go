//コマンドライン引数の連結　（ベンチマーク）
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"time"
)

func main() {
	var s, sep string
	//※１
	now:= time.Now()
	for key, arg := range os.Args[:] {
		s += sep + strconv.Itoa(key) + ":" + arg
		sep = " "
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(s)

	//※２
	now = time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println(time.Now().Sub(now))
}
