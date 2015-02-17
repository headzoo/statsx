package storage

const (
	DefaultContainerMultiplier = 2
)

// NodeContainer is a self expanding container for *Node instances.
type NodeContainer struct {
	Multiplier int64
	nodes []*Node
	length int64
	capacity int64
}

// NewNodeContainer creates and returns a new *NodeContainer instance.
func NewNodeContainer(capacity int64) *NodeContainer {
	return &NodeContainer{
		Multiplier: DefaultContainerMultiplier,
		nodes: make([]*Node, capacity),
		length: 0,
		capacity: capacity,
	}
}

// Length returns the number of nodes in the container.
func (self *NodeContainer) Length() int64 {
	return self.length
}

// Capacity returns the number of nodes the container can currently hold.
func (self *NodeContainer) Capacity() int64 {
	return self.capacity
}

// Nodes returns the nodes in the container.
func (self *NodeContainer) Nodes() []*Node {
	return self.nodes[0:self.length]
}

// Append adds a node to the end of the container.
func (self *NodeContainer) Append(node *Node) {
	if self.length == self.capacity {
		self.expandCapacity()
	}

	self.nodes[self.length] = node
	self.length++
}

// Get returns the node at the specified index.
func (self *NodeContainer) Get(index int64) *Node {
	return self.nodes[index];
}

// expandCapacity expands the container capacity by the multiplier.
func (self *NodeContainer) expandCapacity() {
	self.capacity = self.capacity * self.Multiplier
	nodes := make([]*Node, self.capacity)
	copy(nodes, self.nodes)
	self.nodes = nodes
}
