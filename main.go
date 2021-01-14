package main

import (
  "fmt"
  "@Smashley/Blockchain#blockchain/"



func main() {
  chain := InitBlockChain()

  chain.AddBlock("2nd Block")
  chain.AddBlock("3rd Block")
  chain.AddBlock("4th Block")

  for _, block := range chain.blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data in Block: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Println()

  }
}