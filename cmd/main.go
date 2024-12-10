package main

import (
	"fmt"
	"github.com/kakikubo/blockchain-go/utils"
)

func main() {
	fmt.Println(utils.GetHost())
	//fmt.Println(utils.FindNeighbors("127.0.0.1", 15000, 0, 3, 15000, 15003))
	// Output: [127.0.0.1:15001 127.0.0.1:15002] 自身のポート番号が15000の場合、15001と15002が返される
}
