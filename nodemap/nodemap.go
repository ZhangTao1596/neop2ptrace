package nodemap

import (
	"sync"
)

type NodeMap struct {
	mutex  sync.Mutex
	nodes  Nodes
	matrix Matrix
}

func NewNodeMap() NodeMap {
	return NodeMap{
		nodes:  NewNodes(),
		matrix: NewMatrix(false),
	}
}

func (nm *NodeMap) AddNode(node Node, peers []Node) {
	conns := len(peers)
	if 30 < conns {
		node.group = 4
	} else if 20 < conns {
		node.group = 3
	} else if 10 < conns {
		node.group = 2
	} else {
		node.group = 1
	}
	nm.nodes.AddNode(node)
	for _, v := range peers {
		nm.nodes.AddNode(v)
		nm.matrix.AddEdge(node.ID(), v.ID())
	}
}

func (nm *NodeMap) Count() int {
	return nm.nodes.Count()
}

func (nm *NodeMap) RemoveNode(node string) bool {
	return true
}

func (nm *NodeMap) ToJson() string {
	str := "{"
	str += "\"" + "nodes" + "\"" + ":"
	str += nm.nodes.ToJson()
	str += ","
	str += "\"" + "links" + "\"" + ":"
	str += nm.matrix.ToJson()
	str += "}"
	return str
}

func (nm *NodeMap) String() string {
	return nm.ToJson()
}

func (nm *NodeMap) Start() {

}
