package storage

type NodeType int

const (
	TYPE_DIRECTORY NodeType = 0
	TYPE_FILE NodeType = 1
)

type Node struct {
	Type NodeType
	Name []byte
	Children *NodeContainer
	Counter int64
}

// NewNode creates and returns a *Node instance.
func NewNode(t NodeType, name []byte, child_capacity int64) *Node {
	return &Node{
		Type: t,
		Name: name,
		Children: NewNodeContainer(child_capacity),
	}
}

