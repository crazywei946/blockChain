package main

import (
	"bytes"
	"encoding/gob"
	"crypto/elliptic"
	"io/ioutil"
	"os"
	"fmt"
)

//定义一个wallets.dat文件用来持久化wallets容器
const walletFile  = "wallets.dat"


//定义钱包容器结构体，用来存储密钥对以及匹配的地址
type Wallets struct {
	//map[addr]*wallet

	WalletMap map[string]*Wallet
}

//创建方法生成wallets
func NewWallets()*Wallets  {

	//创建一个空钱包容器  //todo 很绕，还需要理解加强啊！！
	var ws Wallets
	ws.WalletMap=make(map[string]*Wallet)

	//调用loadFile方法将之前的钱包load出来
	ws.LoadFile()

	return &ws
}

//给钱包容器定义方法生成钱包,持久化到本地，并且返回地址
func (ws *Wallets)CreatWallet() string {
	wallet:=NewWallet()  //获得钱包

	addr:=wallet.CreatAddress()  //获得钱包地址

	ws.WalletMap[addr]=wallet  //将新创建的钱包添加到map中，即添加到钱包容器中

	//将钱包容器持久化到本地

	ws.SaveToFile()

	return addr

}

//给钱包容器绑定SaveToFile()的方法
func (ws *Wallets)SaveToFile()  {

	//先将钱包容器编码成字节流
	var buffer bytes.Buffer
	//由于ws中有interface类型，所有需要先注册
	gob.Register(elliptic.P256())
	//创建编码器
	encoder:=gob.NewEncoder(&buffer)

	err:=encoder.Encode(ws)
	if err != nil {
		panic(err)
	}

	//保存字节流
	err=ioutil.WriteFile(walletFile,buffer.Bytes(),0600)
	if err != nil {
		panic(err)
	}

}

//给钱包容器绑定一个loadFile()方法
func (ws *Wallets)LoadFile()  {

	//读去本地文件之前先判断文件是否存在，不存在则创建
	_,err:=os.Stat(walletFile)
	if os.IsNotExist(err) {//表示文件不存在，则创建
		//f,err:=os.Create(walletFile)
		//if err != nil {
		//	panic(err)
		//}
		//
		//defer f.Close()

		return

	}


	//读本地文件
	buf,err:=ioutil.ReadFile(walletFile)
	if err != nil {
		panic(err)
	}

	//将读出来的buf解码成wallets容器

	//注册ws

	gob.Register(elliptic.P256())
	decoder:=gob.NewDecoder(bytes.NewReader(buf))

	var wstemp Wallets
	err=decoder.Decode(&wstemp)
	if err != nil {
		panic(err)
	}

	//将读到的临时文件赋值给ws
	ws.WalletMap=wstemp.WalletMap

}

//给钱包容器绑定方法获得所有的钱包地址
func (ws *Wallets)GetAllAddr()  {
	//遍历容器
	for addr:=range ws.WalletMap{
		fmt.Println(addr)
	}

}


