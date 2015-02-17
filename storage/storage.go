package storage

import "bytes"

const (
	MAX_NAME_LENGTH = 25
)

// PathSeparator is used to separate node names in a path.
var PathSeparator []byte = []byte{'/'}

// WalkerFunc is a callback that may be applied to each node in storage.
type WalkerFunc func(*Node, []byte)

// Storage represents a directory structure of nodes.
type Storage struct {
	// Root is the root directory.
	Root *Node
}

// New creates and returns a *Storage instance.
func New(child_capacity int64) *Storage {
	return &Storage{
		Root: NewNode(TYPE_DIRECTORY, []byte("metrics"), child_capacity),
	}
}

// Walk calls the given function for each node in storage.
func (self *Storage) Walk(f WalkerFunc) {
	parent_path := []byte{}
	f(self.Root, parent_path)
	path := joinPaths(parent_path, self.Root.Name)
	self.walkChildren(f, path, self.Root.Children.Nodes());
}

// walkChildren calls the given function on each of the given nodes, and the
// child nodes recursivly.
func (self *Storage) walkChildren(f WalkerFunc, parent_path []byte, nodes []*Node) {
	for _, node := range nodes {
		f(node, parent_path)
		path := joinPaths(parent_path, node.Name)
		self.walkChildren(f, path, node.Children.Nodes());
	}
}

// joinPaths combines two node paths using the path separator.
func joinPaths(parent, child []byte) []byte {
	return bytes.Join(
		[][]byte{parent, child},
		PathSeparator,
	)
}
