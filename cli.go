package main

import (
	"os"
	"fmt"
	"strconv"
)

//定义一个命令控制器
type CLI struct {
	bc *BlockChain //由于需要操作区块连所以结构体中包含区块连
}

const Usage = `
	addBlock --data DATA "add a block"
    printChain "print block Chain"
	getbance --addr DATA "查询指定地址的余额"
	send from to amount data miner
	newwallet "创建一个新钱包"
	alladdress "获得所有钱包地址"

`

//定义方法对命令行参数进行解析控制执行相应的操作
func (cli *CLI) Run() {

	//获得命令行参数并且解析
	args := os.Args

	if len(args) < 2 {

		//给出提示信息
		fmt.Println(Usage)
		return
	}

	//对第2个参数进行判断

	switch args[1] {
	case "addBlock":
		//表示要继续addBlock
		if len(args) == 4 && args[2] == "--data" {
			//addblock
			cli.AddBlock(args[3])

			fmt.Println("添加区块成功")

		} else {
			fmt.Println("命令行格式为:")
			fmt.Println(Usage)
			return
		}

	case "getbance":
		//判断命令行参数是否符合要求
		if len(args) == 4 && args[2] == "--addr" {
			//执行查询制定地址的余额
			cli.GetBance(args[3])

		} else {
			fmt.Println("命令行格式为:")
			fmt.Println(Usage)
			return
		}

	case "send":
		if len(args) == 7 {
			from := args[2]
			to := args[3]
			amount, _ := strconv.ParseFloat(args[4], 64)
			data := args[5]
			miner := args[6]

			//调用函数进行转账操作
			cli.Send(from, to, amount, data, miner)

			fmt.Println("转账成功")

		} else {
			fmt.Println("转账格式不正确")
			fmt.Println(Usage)
		}

	case "newwallet":
		fmt.Println("创建钱包成功")
		cli.CreatWallet()

	case "alladdress":
		cli.GetAllAddress()

	case "printChain":
		//表示要printChain
		cli.PriteInfo()

	default:
		fmt.Println(Usage)
	}

}
