package router

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// compare func
type Compare func(a, b []byte) int

type RootItem struct {
	Root         *Node
	Compare      Compare
	size, degree uint
}

type Node struct {
	Parent     *Node
	RouterData []*Root
	Children   []*Node
}

type Root struct {
	url []byte
	api uint
}

func NewRouterList(d uint) *RootItem {
	return &RootItem{degree: d, Compare: byteCompare}
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func byteCompare(a, b []byte) int {
	return bytes.Compare(a, b)
}

func (router *RootItem) Put(data []byte, api uint) {
	item := &Root{data, api}

	fmt.Println(item)
}
