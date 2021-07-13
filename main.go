package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("adb", "devices")
	output,err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	i := 0
	for{
		if i > 0 {
			fmt.Println("成功，按任意键退出...")
			keyboard.GetSingleKey()
			os.Exit(1)
		}
		fmt.Println("=======Phigros 数据备份工具=======\n")
		fmt.Println("请输入要执行的操作对应的数字\n")
		fmt.Println("1. 备份")
		fmt.Println("2. 恢复\n")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "1" {
			fmt.Println("请打开手机选择备份")

			res := exec.Command("adb", "backup","-f","PhigrosBackup.ab","com.PigeonGames.Phigros")
			exec.Command("adb", "shell","input","keyevent","82")
			// adb shell input keyevent 82
			// 执行命令，并返回结果
			outputs,err := res.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outputs))
			i++
		}
		if string(char) == "2" {
			fmt.Println("请打开手机选择恢复")
			res := exec.Command("adb", "restore","PhigrosBackup.ab")
			// 执行命令，并返回结果
			outputs,err := res.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outputs))
			i++
		}
	}
}