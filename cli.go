package main

import (
	"os"
	"fmt"
)

//定义一个命令控制器
type CLI struct {
	bc *BlockChain  //由于需要操作区块连所以结构体中包含区块连
}

const Usage  =
	`addBlock --data DATA "add a block"
    printChain "print block Chain"
	getbance --addr DATA "查询指定地址的余额"`


//定义方法对命令行参数进行解析控制执行相应的操作
func (cli *CLI)Run()  {

	//获得命令行参数并且解析
	args:=os.Args

	
	if len(args)<2 {

		//给出提示信息
		fmt.Println(Usage)
		return
	}

	//对第2个参数进行判断

	switch args[1] {
	case "addBlock":
		//表示要继续addBlock
		if len(args)==4 && args[2]=="--data" {
			//addblock
			cli.AddBlock(args[3])

			fmt.Println("添加区块成功")

		}else {
			fmt.Println("命令行格式为:")
			fmt.Println(Usage)
			return
		}

	case "getbance":
		//判断命令行参数是否符合要求
		if len(args)==4 && args[2]=="--addr"{
			//执行查询制定地址的余额
			cli.GetBance(args[3])

		}else {
			fmt.Println("命令行格式为:")
			fmt.Println(Usage)
			return
		}

	case "printChain":
		//表示要printChain
		cli.PriteInfo()

	default:
		fmt.Println(Usage)
	}

}

