package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("adb", "devices")
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	for{
		var inputs int
		fmt.Println("==================")
		fmt.Println("请输入要执行的操作对应的数字")
		fmt.Println("1. 备份")
		fmt.Println("2. 恢复")
		fmt.Scanf("%d\n", &inputs)
		if inputs == 1 {
			fmt.Println("请打开手机选择备份")
			res := exec.Command("adb", "backup","-f","PhigrosBackup.ab","com.PigeonGames.Phigros")
			// 执行命令，并返回结果
			outputs,err := res.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outputs))
		}
		if inputs == 2 {
			fmt.Println("请打开手机选择恢复")
			res := exec.Command("adb", "restore","PhigrosBackup.ab")
			// 执行命令，并返回结果
			outputs,err := res.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outputs))
		}
	}
}