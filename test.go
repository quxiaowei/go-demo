package main

import (
	"fmt"

	"github.com/quxiaowei/go-demo/examples"
)

func main() {
	var exp string
	fmt.Printf("Enter Example Name:")
	fmt.Scanln(&exp)

	switch exp {
	case "structs":
		examples.Structs()
	case "switch":
		examples.Switch()
	case "array":
		examples.Array()
	case "slice":
		examples.Slice()
	case "map":
		examples.Map()
	case "range":
		examples.Range()
	case "sum":
		examples.Sum()
	case "pointer":
		examples.Pointer()
	case "channel":
		examples.ChannelTest()
	case "channel_select":
		examples.ChannelSelect()
	case "channel_range":
		examples.ChannelRange()
	case "steamline":
		examples.Steamline(10, 20)
	default:
		fmt.Printf("no example named \"%s\"\n", exp)
		println("...")
		main()
	}
}
