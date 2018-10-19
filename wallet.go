package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
)

//创建一个结构体用来保存秘钥对
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey  //私钥

	PulickKey  []byte  //使用公钥中的x y的拼接字节作为公钥，便于后边的传参

}

//定义方法，生成结构提对象
func NewWallet() *Wallet {

	//使用ecdsa生成秘钥对
	privateKey,err:=ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	if err != nil {
		panic(err)
	}

	//获得公钥
	public:=privateKey.PublicKey

	//将公钥中的xy拼接成字节流
	xbyte:=public.X.Bytes()
	ybyte:=public.Y.Bytes()

	publicKey:=append(xbyte, ybyte...)

	wallet:=Wallet{privateKey,publicKey}

	return &wallet

}

//创建函数对公钥进行hash
func HashPublick(publick []byte) []byte {

	hash256:=sha256.Sum256(publick)

	//进行160hash
	//创建160hash指针
	hash160:=ripemd160.New()
	_,err:=hash160.Write(hash256[:])
	if err != nil {
		panic(err)
	}

	//进行hsah
	publikHash:=hash160.Sum(nil)

	return publikHash


}

//给wallet绑定方法生成对应的地址
func (ws *Wallet)CreatAddress() string {

	//通过wallet获得公钥
	publicKey:=ws.PulickKey

	//调用函数//通过公钥进行hash，然后160hsah
	publicHash:=HashPublick(publicKey) //字节为20的公钥hash

	//将hash与版本号进行拼接得到长度21位的串
	payload:= append([]byte{version}, publicHash...)

	//对21位的串进行2次256hash，取前四位得到checknum
	hashFirst:=sha256.Sum256(payload)
	hashSecond:=sha256.Sum256(hashFirst[:])

	checkSum:=hashSecond[:4]

	//将21+4得到的25串进行base58得到地址
	hashTemp:=append(payload, checkSum...)

	address:=base58.Encode(hashTemp)

	return address
}
