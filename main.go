package main

func main() {

	//创建区块链
	bc := NewBlcokChain()

	cli:=CLI{bc}

	cli.Run()
	//
	//bc.AddBlock("小王赚了100BTC给给你春花")
	//bc.AddBlock("小王赚了100BTC给给你笑话")

	//遍历迭代器
	//1.获得迭代器
	//it := bc.NewIterator()


}
