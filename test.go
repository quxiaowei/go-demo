package main

import (
	"github.com/quxiaowei/go-demo/examples"
)

func main() {
	exp := "channel_select"

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
	}
}
