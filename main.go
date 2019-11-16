package main

import (
	LennyPkgSpt "./src/LennyPackagingScript"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("请输入: 参数!")
		fmt.Println("----->. ", "-s 打单个的包")
		fmt.Println("----->.", "-from 从指定盘口开始全部打完.")
		fmt.Println("----->.", "-from * -to * 从指定盘口打到指定盘口")
		os.Exit(0)
	}else if args[0] == "-s" {		// 打单个的包
		targetName := args[1]
		LennyPkgSpt.PackagingTarget(targetName)
	}else if args[0] == "-from" && args[1] != "-to" {   	//从某一个盘口开始全部打完
		targetName := args[1]
		LennyPkgSpt.PackagingFrom(targetName)
	}else if args[0] == "-from" && args[1] == "-to" { 		//从指定盘口开始打到 指定盘口

	}
}
