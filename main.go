package main

import (
	"fmt"
	"github.com/zhangyiming748/countYTB/opt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("第二个参数需要填写当前的速度(kbps)")
		os.Exit(0)
	}
	Speed := os.Args[1]
	f, _ := strconv.ParseFloat(Speed, 64)
	speed := opt.Count(f)
	fmt.Printf("real speed is %.3fMB", speed)
}
