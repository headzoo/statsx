package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeContainer(t *testing.T) {
	nc := NewNodeContainer(10)
	
	assert.Equal(t, 0, nc.Length())
	assert.Equal(t, 10, nc.Capacity())
}

func TestNodeContainerAppend(t *testing.T) {
	nc := NewNodeContainer(2)
	
	nc.Append(NewNode(TYPE_DIRECTORY, []byte{}, 2))
	assert.Equal(t, 1, nc.Length())

	nc.Append(NewNode(TYPE_DIRECTORY, []byte{}, 2))
	assert.Equal(t, 2, nc.Length())

	nc.Append(NewNode(TYPE_DIRECTORY, []byte{}, 2))
	assert.Equal(t, 3, nc.Length())
	
	assert.Equal(t, (2 * DefaultContainerMultiplier), nc.Capacity())
}

func TestNodeContainerGet(t *testing.T) {
	node0 := NewNode(TYPE_DIRECTORY, []byte{}, 2)
	node1 := NewNode(TYPE_DIRECTORY, []byte{}, 2)
	nc := NewNodeContainer(2)
	nc.Append(node0)
	nc.Append(node1)
	
	assert.Equal(t, node0, nc.Get(0))
	assert.Equal(t, node1, nc.Get(1))
}
