package main

import "fmt"

type Block struct {
	ID    int
	Data  string
	Color string
}

var Blocks []Block

func DisplayAllBlocks() {
	fmt.Println("Displaying all blocks:")
	for _, block := range Blocks {
		fmt.Printf("Block ID: %d, Data: %s, Color: %s\n", block.ID, block.Data, block.Color)
	}
	fmt.Println()
}

func NewBlock(id int, data string, color string) {
	newBlock := Block{
		ID:    id,
		Data:  data,
		Color: color,
	}
	Blocks = append(Blocks, newBlock)
	fmt.Printf("New block created - ID: %d, Data: %s, Color: %s\n", id, data, color)
}

func ModifyBlock(id int, newData string) {
	for i, block := range Blocks {
		if block.ID == id {
			Blocks[i].Data = newData
			fmt.Printf("Block with ID %d modified - New Data: %s\n", id, newData)
			return
		}
	}
	fmt.Printf("Block with ID %d not found\n", id)
}

func main() {
	NewBlock(1, "Data1", "Red")
	NewBlock(2, "Data2", "Blue")
	DisplayAllBlocks()

	ModifyBlock(1, "Updated Data 1")
	DisplayAllBlocks()
}
