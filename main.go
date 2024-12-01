package main

import (
	"fmt"
	"github.com/kakikubo/blockchain-go/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	//fmt.Println(w.PrivateKey())
	//fmt.Println(w.PublicKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
