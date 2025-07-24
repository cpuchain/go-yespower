package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cpuchain/go-yespower"
)

func main() {
	input := "eebb7bf9a8c813b5e0a03ce627bd1a0c836e0a89793743666dc82b83e28e8f00"

	in, _ := hex.DecodeString(input)

	out := hex.EncodeToString(yespower.Hash(in, uint32(2048), uint32(32), ""))

	out2 := hex.EncodeToString(yespower.YespowerNative(in, 2048, 32, ""))
	
	fmt.Println(input, out)

	fmt.Println(input, out2)
}