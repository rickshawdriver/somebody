package router

import (
	"bytes"
)

// compare func
type Compare func(a, b []byte) int

type RouterDegree int

type RootItem struct {
	Root         *Node
	Compare      Compare
	size, degree int
}

type Node struct {
	Parent     *Node
	RouterData []*Item
	Children   []*Node
}

type Item struct {
	url []byte
	api uint
}

func NewRouterList(d int) *RootItem {
	return &RootItem{degree: d, Compare: byteCompare}
}

func byteCompare(a, b []byte) int {
	return bytes.Compare(a, b)
}

func (router *RootItem) Put(data []byte, api uint) {
	item := &Item{data, api}

	if router.Root == nil {
		router.Root = &Node{Children: []*Node{}, RouterData: []*Item{item}}
	}

	if router.insert(router.Root, item) {
		return
	}

	router.size++
}

func (router *RootItem) insert(node *Node, item *Item) bool {
	if router.isLeaf(node) {
		return router.insertIntoLeaf(node, item)
	}
	return router.insertIntoInternal(node, item)
}

func (router *RootItem) insertIntoInternal(node *Node, item *Item) bool {
	insertPosition, found := router.search(node, item.url)
	if found {
		node.RouterData[insertPosition] = item
		return false
	}
	return router.insert(node.Children[insertPosition], item)
}

func (router *RootItem) insertIntoLeaf(node *Node, item *Item) bool {
	insertPosition, found := router.search(node, item.url)
	if found {
		node.RouterData[insertPosition] = item
		return false
	}
	node.RouterData = append(node.RouterData, nil)
	copy(node.RouterData[insertPosition+1:], node.RouterData[insertPosition:])
	node.RouterData[insertPosition] = item
	router.split(node)
	return true
}

func (router *RootItem) search(node *Node, key []byte) (index int, found bool) {
	low, high := 0, len(node.RouterData)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := router.Compare(key, node.RouterData[mid].url)
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

func (router *RootItem) isLeaf(node *Node) bool {
	return len(node.Children) == 0
}

func (router *RootItem) split(node *Node) {
	if len(node.Children) != router.degree {
		return
	}

	if node == router.Root {
		router.splitRoot()
		return
	}

	router.splitNonRoot(node)
}

func (router *RootItem) middle() int {
	return (router.degree - 1) / 2
}

func (router *RootItem) splitRoot() {
	middle := router.middle()

	left := &Node{RouterData: append([]*Item(nil), router.Root.RouterData[:middle]...)}
	right := &Node{RouterData: append([]*Item(nil), router.Root.RouterData[middle+1:]...)}

	if !router.isLeaf(router.Root) {
		left.Children = append([]*Node(nil), router.Root.Children[:middle+1]...)
		right.Children = append([]*Node(nil), router.Root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	newRoot := &Node{
		RouterData: []*Item{router.Root.RouterData[middle]},
		Children:   []*Node{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	router.Root = newRoot
}

func setParent(nodes []*Node, parent *Node) {
	for _, node := range nodes {
		node.Parent = parent
	}
}

func (router *RootItem) splitNonRoot(node *Node) {
	middle := router.middle()
	parent := node.Parent

	left := &Node{RouterData: append([]*Item(nil), node.RouterData[:middle]...), Parent: parent}
	right := &Node{RouterData: append([]*Item(nil), node.RouterData[middle+1:]...), Parent: parent}

	if !router.isLeaf(node) {
		left.Children = append([]*Node(nil), node.Children[:middle+1]...)
		right.Children = append([]*Node(nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := router.search(parent, node.RouterData[middle].url)

	parent.RouterData = append(parent.RouterData, nil)
	copy(parent.RouterData[insertPosition+1:], parent.RouterData[insertPosition:])
	parent.RouterData[insertPosition] = node.RouterData[middle]

	parent.Children[insertPosition] = left

	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	router.split(parent)
}

func (router RootItem) Get(url []byte) (value interface{}, found bool) {
	node, index, found := router.searchRecursively(router.Root, url)
	if found {
		return node.RouterData[index].api, true
	}
	return nil, false
}

func (router *RootItem) Empty() bool {
	return router.size == 0
}

func (router *RootItem) searchRecursively(startNode *Node, url []byte) (node *Node, index int, found bool) {
	if router.Empty() {
		return nil, -1, false
	}
	node = startNode
	for {
		index, found = router.search(node, url)
		if found {
			return node, index, true
		}
		if router.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}
