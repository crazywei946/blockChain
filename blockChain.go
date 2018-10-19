package main

import (
	"github.com/boltdb/bolt"
	"fmt"
)

//定义一个区块链结构体
type BlockChain struct {
	db *bolt.DB //数据库存储

	tail []byte //最后一个区块当前区块的hash
}

const genensisInfo = "This is GenensisInfo Block This is GenensisInfo Block" //创始区块内容
const blockChainDb = "blockChain.db"                                         //数据库名
const blockChainForm = "blockChainForm"                                      //表明
//const lastHash  ="lastHash" //最后区块的hash的key值

//创建一个生成区块链的方法
func NewBlcokChain() *BlockChain {

	var bc *BlockChain
	//var db *bolt.DB
	var lastHash []byte
	//打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		panic("打开数据库失败")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if err != nil {
			panic("数据库Updata失败")
		}
		//打开数据库中的表，没有则创建表
		bucket := tx.Bucket([]byte(blockChainForm))
		if bucket == nil {
			//表不存在创建表
			bucket, err = tx.CreateBucket([]byte(blockChainForm))
			if err != nil {
				panic("创建表失败")
			}

			//往表中添加入区块
			//1.创始区块的hash作为key,value为区块
			gensisBlock := GenesisBlock(genensisInfo, "老王")
			bucket.Put(gensisBlock.Hash, gensisBlock.Serializa())

			fmt.Println("************************")
			//2.用常量l最为key，value为lastBlockHash
			bucket.Put([]byte("l"), gensisBlock.Hash)

			//3.将最后的hash复制给lasthash
			lastHash = gensisBlock.Hash
			////实列化BlockChain
			//bc = &BlockChain{db, lastHash}
		} else {
			//打开表取出表中数据
			lastHash = bucket.Get([]byte("l"))

		}

		return nil
	})

	//实列化BlockChain
	bc = &BlockChain{db, lastHash}
	return bc

}

//创建一个方法给blockchian中添加区块
func (bc *BlockChain) AddBlock(txs []*Tranction) {

	//获得最后一个区块的hash
	preHash := bc.tail

	//创建新区块
	block := NewBlock(txs, preHash)

	//打开数据库中的表
	//更新数据库表中最后dehash值
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockChainForm)) //dakai 表
		if bucket == nil {
			panic("打开表失败")
		}

		//gengxin 表中数据
		bucket.Put(block.Hash, block.Serializa())
		bucket.Put([]byte("l"), block.Hash)

		//根新连数据
		bc.tail = block.Hash
		return nil
	})

}

//给blockChain绑定方法查询指定地址的可用交易余额
func (bc *BlockChain) FindUTXO(addr string) ([]TXOutput) {

	//遍历所有区块，拿到所有的交易
	//拿到迭代器进行遍历
	it := bc.NewIterator()
	var block *Block

	var UTXO []TXOutput //未花费交易的集合

	spentUTXO := make(map[string][]int64) //已经花费的交易集合

	for {
		block = it.Next()
		//遍历所有交易，拿到所有的outputs
		for _, tx := range block.Trancations {

		OUTPUT:
		//遍历outputs，对又有outputs进行判断
			for i, out := range tx.TXOutputs {
				//1.判断该output是否已经被花费掉
				//看该out是否在所有的txinput中，如果在则表示已经被花费
				//遍历已经花费的交易
				if spentUTXO[string(tx.TXID)] != nil {
					//如果当前交易的id在已花费的交易集合中，则表示需要对当前交易中output进行判断是否存在
					//如果当前交易的id不在已花费的交易集合中，则跳过，不对当前交易中的output进行判断
					for _, indexArry := range spentUTXO[string(tx.TXID)] {

						//找到当前交易id所对应的map时，遍历map的值，判断output的索引是否和map中的值一致，一致则表示此output已经花费掉
						if int64(i) == indexArry {
							//相等则进行下一次循环
							continue OUTPUT
						}

					}
				}

				//2.output未被花费掉，判断该output是否属于指定的地址
				if out.PubkyeHash == addr {
					//3.将符合条件的output放入集合中返回
					UTXO = append(UTXO, out)
				}

			}

			//判断该是不是创世交易，如果不是才进行遍历
			if !tx.IscoinBase() {
				//遍历交易拿到地址对应的所有的input，并加入到一个map[交易ID][]int64集合中
				//key为input的TXid，value为input的index
				for _, input := range tx.TXInputs {
					if input.Sig == addr {
						spentUTXO[string(input.Txid)] = append(spentUTXO[string(input.Txid)], input.Index)

					}
				}

			}

		}

		if len(block.PreHash) == 0 {
			fmt.Println("区块遍历完成")
			break
		}

	}

	return UTXO

}

//给blockChain绑定方法查询制定地址所需要的交易余额，并且返回余额总数
func (bc *BlockChain) FindNeedUTXO(addr string, amount float64) (map[string][]int64, float64) {

	//找到所需要的UTXO
	needUTXO := make(map[string][]int64)

	spentInputs := make(map[string][]int64)
	var total float64

	//遍历区块，获得所有的交易
	//1.获得迭代起进行迭代
	it := bc.NewIterator()
	var block *Block

	for {
		block = it.Next()
		//遍历所有交易，获得outputs
		for _, tx := range block.Trancations {

		OUTPUT:
		//遍历outputs，获得制定地址的所有output
			for outindex, out := range tx.TXOutputs {
				//判断当前交易中是否含有未花费的outputs
				if spentInputs[string(tx.TXID)] != nil {
					//判断out是否已经被花费过
					for _, index := range spentInputs[string(tx.TXID)] {
						if int64(outindex) == index {
							continue OUTPUT
						}
					}
				}

				//判断out是否属于指定的地址
				if out.PubkyeHash == addr {

					if total < amount {

						//属于则将这个output加入到needUTXO中
						needUTXO[string(tx.TXID)] = append(needUTXO[string(tx.TXID)], int64(outindex))

						//判断查询到的value是否不小于需要转帐的amount
						total += out.Value
					}
					if total >= amount {
						//结束遍历，直接返回needutxo和total
						return needUTXO, total

					}
				}
			}

			//因为创世区块没有inputs，所以先判断是否是创世区块
			if !tx.IscoinBase() {

				//而找到指定地址所有的inputs，添加到一个固定集合中map[TXID][index]
				for _, input := range tx.TXInputs {

					if input.Sig == addr {
						spentInputs[string(input.Txid)] = append(spentInputs[string(input.Txid)], input.Index)
					}

				}

			}

		}

		//判断是否已经对区块遍历完成
		if len(block.PreHash) == 0 {
			break
		}

	}
	//对output进行判断，看是否已经被消费过
	//拿到output中的value与amount进行判断，钱不够则将outout加入到needutxo中

	return needUTXO, total

}

//对UTXO代码进行优化
//抽取一个公用函数，找到可用的outputs所对应的trancations
//func (bc *BlockChain) FindUTXOTrancations(addr string) []*Tranction {
//
//	//遍历区块
//	it := bc.NewIterator()
//
//	var UTXOTrancastions []*Tranction
//
//	spendUTXO := make(map[string][]int64)
//
//	for {
//		block := it.Next()
//		//遍历交易
//		for _, tx := range block.Trancations {
//
//			OUTPUT:
//			//遍历outputs
//			for i, output := range tx.TXOutputs {
//
//				//判断输出是否有效
//				//TODO
//				if spendUTXO[string(tx.TXID)]!=nil {
//
//					for _,index:=range spendUTXO[string(tx.TXID)]{
//						//如果当前output的索引值在spend中，则表示此output已经被花费掉，继续下一个output判断
//						if int64(i)==index {
//							continue OUTPUT
//						}
//
//					}
//
//				}
//
//				//判断输出是否属于制定地址
//				if output.PubkyeHash == addr {
//
//					//将此地址属于的tx加入到交易数组集合中
//					UTXOTrancastions = append(UTXOTrancastions, tx)
//
//				}
//
//			}
//
//			//判断是否是创世区块，不是则不需要判断txinput
//			if !tx.IscoinBase() {
//
//				//遍历inputs，将所有inputs放入到一个map中map[txid][]int64
//				for _, input := range tx.TXInputs {
//
//					spendUTXO[string(input.Txid)] = append(spendUTXO[string(input.Txid)], input.Index)
//
//				}
//			}
//
//		}
//
//		if len(block.PreHash) == 0 {
//			//区块遍历完成
//			break
//		}
//	}
//
//	return UTXOTrancastions
//
//}
