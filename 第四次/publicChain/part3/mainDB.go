package main

import (
	"publicChain/part3/BLC"
)

func main()  {

	// 创世区块

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	 //新区块
	blockchain.AddBlockToBlockchain("send 100RMB a")

	blockchain.AddBlockToBlockchain("send 200RMB To b")

	blockchain.AddBlockToBlockchain("send 300RMB To c")

	blockchain.AddBlockToBlockchain("send 50RMB To d")

	blockchain.Printchain()

}
